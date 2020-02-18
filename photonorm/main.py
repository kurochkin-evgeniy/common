#!/usr/bin/python
# -*- coding: UTF-8 -*-

try:
    import pyexiv2
except:
    print ("Run: pip install pyexiv2")
    sys.exit(1)


import sys
import os
import datetime
import remove_dubles
import organize_folders
import print_info
import add_tags
import change_exif_time


def PerformInfo(path, all):
    print_info.PrintExifInfo(path, all)


def PerformRemoveDublicatedFiles(path, emulate):
    print(remove_dubles.RemoveDublicatedFiles(path, emulate))


def PerformModifyExifTime(path):
    filelist = os.listdir(path)
    for filename in filelist:
        try:
            print ("Process file: ", filename)
            change_exif_time.ModifyImageFile(path + '\\' + filename)
        except Exception as e:
            print (e)


def PerformOrganize(path):
    organize_folders.OrganizeFolderStructure(path)


def PerformTags(path):
    add_tags.AddTags(path)


def PerformMake(path):
    PerformRemoveDublicatedFiles(path, False)
    PerformOrganize(path)


if __name__ == '__main__':
    command = sys.argv[1]
    path = sys.argv[2]

    if command == 'dub':
        PerformRemoveDublicatedFiles(path, True)
    elif command == 'time':
        PerformModifyExifTime(path)
    elif command == 'info':
        PerformInfo(path, False)
    elif command == 'infoall':
        PerformInfo(path, True)
    elif command == 'org':
        PerformOrganize(path)
    elif command == 'tags':
        PerformTags(path)
    elif command == 'make':
        PerformMake(path)
