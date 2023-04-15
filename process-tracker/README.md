# Process Tracker

Go version of a script I used with my operating systems class to show a message
across all my terminals when a process was using too much CPU or memory or was running for
too long.

### Read before running

1. You need to be in a login shell to receive broadcast messages. Broadcasts can only be
   sent to logged-in users
    1. To switch to log in shell, simply ssh into your pc like: `ssh $HOSTNAME`
    2. To verify you're in a login shell, run `echo $0`. If the first character is a dash,
       you're in login shell.
2. There's a **_very slim_** chance sending broadcasts will fail if the parameters are too loose.
   i.e. if you use cmd line args that will print out almost all the processes. E.g
   `./proc-tracker -cpu 0.2 -time 2s`

3. Sending notifications and broadcasts won't work on windows and I genuinely do not care.
   This is a linux supremacy zone
    1. It should work on Mac if `notify-send` and `wall` are installed
4. If the cpu or mem usage % could not be gotten for some reason, it will default to `999`
    1. If the time could not be gotten, it will default to something around `999:10:00`
    2. If the process name or owner could not be gotten, it will default to `--------"

### Precompiled binaries

1. See release section.

### Usage

For boolean flags: `-flag=true`. The following are also valid: `1, 0, t, f, T, F, true, false,
TRUE, FALSE, True, False`

See `process-tracker -h` for usage

      Usage of ./proc-tracker:
         -bcst
            To enable or disable broadcast messages (default true)
         -cpu float
            Max CPU usage % to trigger notification (default 70)
         -gui
            To enable or disable GUI notifications. 'notify-send' is needed (default true)
         -i duration
            Interval between each check (default 2m0s)
         -mem float
            Max memory usage % to trigger notification (default 70)
         -o string
            Sort order. 'asc' or 'dsc' (default "dsc")
         -s string
            Which column to sort by (default "CPU %")
         -term
            Show messages in the same terminal the program is running in (default false)
         -time duration
            Max time a process can run before triggering notification (default 1h0m0s)

### Build and compile yourself because you don't trust me

1. The code was written with go version 1.20.3. So, you need to have that version or higher
    1. I didn't really look much into the versions so it might work with older versions
2. See Makefile, and note your platform and architecture. Try `amd64` version first
3. run `make <proc-tracker-platform-arch>`
4. or you can simply run `make` and try everything to find which works
