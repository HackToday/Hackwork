#define _GNU_SOURCE
#include <sys/types.h>
#include <sys/wait.h>
#include <sys/mount.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sched.h>
#include <signal.h>
#include <unistd.h>
#include <limits.h>
#include <fcntl.h>
#include <errno.h>


#define STACK_SIZE (1024 * 1024)

// sync primitive

static char child_stack[STACK_SIZE];
struct  child_args {
  char **argv;
  int pipe_fd[2];
  
};

int child_main(void* arg)
{
  char c;
  struct child_args *args = (struct child_args *) arg;

  // init sync primitive
  close(args->pipe_fd[1]);

  // setup hostname
  //printf(" - [%5d] World !\n", getpid());
  sethostname("In-Namespace", 12);

  // remount "/proc" to get accurate "top" && "ps" output
  //mount("proc", "/proc", "proc", 0, NULL);
  // wait...
  read(args->pipe_fd[0], &c, 1);

  execvp(args->argv[0], args->argv);
  exit(EXIT_FAILURE);
}

static void
update_map(char *mapping, char *map_file)
{
    int fd, j;
    size_t map_len;     /* Length of 'mapping' */

    /* Replace commas in mapping string with newlines */
    map_len = strlen(mapping);
    for (j = 0; j < map_len; j++)
        if (mapping[j] == ',')
            mapping[j] = '\n';

    fd = open(map_file, O_RDWR);
    if (fd == -1) {
        fprintf(stderr, "open %s: %s\n", map_file, strerror(errno));
        exit(EXIT_FAILURE);
    }

    if (write(fd, mapping, map_len) != map_len) {
        fprintf(stderr, "write %s: %s\n", map_file, strerror(errno));
        exit(EXIT_FAILURE);
    }

    close(fd);
}

int main(int argc, char *argv[])
{
  // init sync primitive
  struct child_args args;
  char *uid_map, *gid_map;
  int opt;
  int flags = 0;
  char map_path[PATH_MAX];
  while ((opt = getopt(argc, argv, "+imnpuUM:G:")) != -1) {
      switch (opt) {
      case 'i': flags |= CLONE_NEWIPC;        break;
      case 'm': flags |= CLONE_NEWNS;         break;
      case 'n': flags |= CLONE_NEWNET;        break;
      case 'p': flags |= CLONE_NEWPID;        break;
      case 'u': flags |= CLONE_NEWUTS;        break;
      case 'M': uid_map = optarg;             break;
      case 'G': gid_map = optarg;             break;
      case 'U': flags |= CLONE_NEWUSER;       break;
      default:  printf("Nothing helper here.");
      }
   }
  args.argv = &argv[optind];
  if (pipe(args.pipe_fd) == -1) {
      perror("pipe");
      exit(EXIT_FAILURE);
  }
  printf(" - [%5d] Hello ?\n", getpid());

  pid_t child_pid = clone(child_main, child_stack+STACK_SIZE,
      flags | SIGCHLD, &args);

  if (child_pid == -1) {
      exit(EXIT_FAILURE);
  }

  if (uid_map != NULL) {
      snprintf(map_path, PATH_MAX, "/proc/%ld/uid_map", (long)child_pid);
      update_map(uid_map, map_path);
  }
  if (gid_map != NULL) {
      snprintf(map_path, PATH_MAX, "/proc/%ld/setgroups", (long)child_pid);
      update_map("deny", map_path);
      snprintf(map_path, PATH_MAX, "/proc/%ld/gid_map", (long)child_pid);
      update_map(gid_map, map_path);
  }
  // further init here (nothing yet)

  // signal "done"
  close(args.pipe_fd[1]);

  waitpid(child_pid, NULL, 0);
  
  exit(EXIT_FAILURE);
}
