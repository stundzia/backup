package main

import (
	"backup/dirs"
	"time"
)

const dirToBackup = "/home/stundzia/go/src/backup/target"
const dirBackup = "/home/stundzia/go/src/backup/dest"
//const dirToBackup = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\ATLAS\\ShooterGame\\Saved\\SavedAtlasLocal"
//const dirBackup = "C:\\Users\\pauli\\OneDrive\\Stalinis kompiuteris\\atlas\\backup"

func main() {
	m := dirs.SetupBackupManager(dirToBackup, dirBackup, 4)
	m.Backup()
	// Time to read any errors in windows.
	time.Sleep(2 * time.Second)
}