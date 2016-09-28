Common Configuration Issue
==============================

## Disable Password Authentication

1. Edit `/etc/ssh/sshd_config` with

  ```
# Change to no to disable tunnelled clear text passwords
PasswordAuthentication no
```

2. restart ssh service

  ```
  systemctl restart sshd
  ```
