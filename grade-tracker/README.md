I could have done all this with awk in less than 6 lines smh. Just trying to get conformable with Go's
syntax.

# Grade Tracker

Program to simply report the current grade in a course. Most of similar tools like this online don't have
any way to save the grades.

## Usage

`./tracker-{os-arch} <path_to_file>`

The file:

1. Needs to be in this format: `Name, achieved, possible, % of final grade` (csv, see `sample.bighead`)
2. File extension does not matter
3. `#` is treated like a comment and ignored
4. See `result.bighead` for example output
5. Keep 'name' under 15 characters to preserve formatting
6. If there's any issue with the file, execution stops immediately

## If you 100% trust me and want to use the binaries I compiled

1. You don't need to install go
2. Download your version from the `release` section on the repo homepage (if I can figure it out).
   1. Try the amd64 version first
3. There's no gui, so run from command line

## If you hate me and want to compile it yourself (racially motivated action)

1. Run the `build-all.sh` script or
2. Compile for 64-bit linux: `env GOOS=linux GOARCH=amd64 go build -o tracker-linux-amd64 tracker.go`
3. Compile for 64-bit windows: `env GOOS=windows GOARCH=amd64 go build -o tracker-win-amd64.exe tracker.go`
4. Compile for 64-bit mac: `env GOOS=darwin GOARCH=amd64 go build -o tracker-mac-amd64 tracker.go`
5. Go obviously needs to be installed**
6. If the generated binary doesn't work, you're probably on arm65. Change the value of GOARCH to arm64**
