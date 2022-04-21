# elf

```
import "debug/elf"

Package elf implements access to ELF object files.
```




## func NewFile(r io.ReaderAt) (*File, error)

```
NewFile creates a new File for accessing an ELF binary in an underlying reader.
The ELF binary is expected to start at position 0 in the ReaderAt.
```



---
