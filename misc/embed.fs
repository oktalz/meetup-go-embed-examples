func (f FS) Open(name string) (fs.File, error)
func (f FS) ReadFile(name string) ([]byte, error)
func (f FS) ReadDir(name string) ([]fs.DirEntry, error)
