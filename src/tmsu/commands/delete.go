package main

import (
           "errors"
       )

type DeleteCommand struct {}

func (this DeleteCommand) Name() string {
    return "delete"
}

func (this DeleteCommand) Summary() string {
    return "deletes one or more tags"
}

func (this DeleteCommand) Help() string {
    return `tmsu delete TAG...

Permanently deletes the tag(s) specified.`
}

func (this DeleteCommand) Exec(args []string) error {
    if len(args) == 0 { return errors.New("No tags to delete specified.") }

    db, error := OpenDatabase(databasePath())
    if error != nil { return error }
    defer db.Close()

    for _, tagName := range args {
        error = this.deleteTag(db, tagName)
        if error != nil { return error }
    }

    return nil
}

func (this DeleteCommand) deleteTag(db *Database, tagName string) error {
    tag, error := db.TagByName(tagName)
    if error != nil { return error }

    error = db.RemoveFileTagsByTagId(tag.Id)
    if error != nil { return error }

    error = db.DeleteTag(tag.Id)
    if error != nil { return error }

    return nil
}