dirEntries, err := fs.ReadDir("embedded-folder")
if err != nil {
	log.Fatal(err)
}
for _, entry := range dirEntries {
	if entry.IsDir() {
		...
		continue
	}
	fmt.Println(path.Join("embedded-folder", newName(entry.Name())))
}
