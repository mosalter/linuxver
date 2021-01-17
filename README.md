# linuxver
--
    import "github.com/mosalter/linuxver"

Package linuxver parses and compares upstream linux kernel version strings.

## Documentation

```go
var (
	// NoVersion is a special linux version for commits which don't precede
	// a version tag (those following newest version tag when HEAD is not
	// itself tagged with a version).
	NoVersion = &LinuxVersion{255, 0, 0, 0}
)
```

#### type LinuxVersion

```go
type LinuxVersion struct {
	// Major holds the major version number. A major value of 255 is special
	// and indicates "no version". This is useful when the linux kernel tree
	// has commits at HEAD which do not precede a version tag (HEAD itself
	// is not tagged with a version).
	Major uint8
	// Minor holds the minor version number.
	Minor uint8
	// Rel holds a release version number (only used in v2.6 kernels).
	Rel uint8
	// RC holds the release candidate number or zero for the final release.
	RC uint8
}
```

LinuxVersion represents an upstream Linux version.

#### func  New

```go
func New(v string) *LinuxVersion
```
New tries to convert a given string into a LinuxVersion. If successful, a
pointer to the LinuxVersion is returned, else nil is returned.

#### func (*LinuxVersion) ComesAfter

```go
func (v1 *LinuxVersion) ComesAfter(v2 *LinuxVersion) bool
```
ComesAfter returns true if v1 is newer than v2

#### func (*LinuxVersion) ComesBefore

```go
func (v1 *LinuxVersion) ComesBefore(v2 *LinuxVersion) bool
```
ComesBefore returns true if v1 is older than v2

#### func (*LinuxVersion) Equals

```go
func (v1 *LinuxVersion) Equals(v2 *LinuxVersion) bool
```
Equals returns true if LinuxVersion v1 == v2

#### func (*LinuxVersion) String

```go
func (v *LinuxVersion) String() string
```
String converts the LinuxVersion to a readable string

## Credits

 * [Mark Salter](https://github.com/mosalter)

## License

The GPL v3 License (GPLv3) - see [`LICENSE.md`](https://github.com/mosalter/linuxver/blob/master/LICENSE.md) for more details

