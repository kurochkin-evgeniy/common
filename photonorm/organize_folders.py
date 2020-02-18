#!/usr/bin/python
# -*- coding: UTF-8 -*-

import sys
import os
import datetime
import shutil
import pyexiv2


def IsImageFile(path):
    filename, ext = os.path.splitext(path)
    ext.upper()
    if ext == '.NEF' or ext == '.JPG' or ext == '.JPEG':
        return True
    return False


def GetShotDate(path):
    meta = pyexiv2.ImageMetadata(path)
    meta.read()
    if "Exif.Photo.DateTimeOriginal" in meta.exif_keys:
        return meta['Exif.Photo.DateTimeOriginal'].value

    return None


def GetAllFilesWithShotDate(path):
    resultFiles = []
    for root, dirs, files in os.walk(path):
        for _file in files:
            fullPath = os.path.join(root, _file)
            if IsImageFile(fullPath):
                resultFiles.append((fullPath, GetShotDate(fullPath)))
    return resultFiles


def OrganizeFolderStructure(path):
    p = os.path.split(path)
    newPath = os.path.join(p[0], p[1] + '_')
    shutil.rmtree(newPath)
    os.mkdir(newPath)

    photos = GetAllFilesWithShotDate(path)

    if len(photos) == 0:
        return

    photos.sort(key=lambda x: x[1])

    currentDate = photos[0][1]
    currentFolderName = os.path.join(newPath, currentDate.strftime('%d %B %y'))
    os.mkdir(currentFolderName)
    index = 1
    dirIndex = 1
    for file in photos:
        if currentDate.date() != file[1].date():
            currentDate = file[1]
            currentFolderName = os.path.join(newPath, '{}. {}'.format(
                dirIndex, currentDate.strftime('%d %B %y')))
            os.mkdir(currentFolderName)
            dirIndex += 1

        filename, file_extension = os.path.splitext(file[0])
        file_extension.upper()
        p = 'DSC_{:0>4}{}'.format(index, file_extension)
        newPhotoPath = os.path.join(currentFolderName, p)
        print ('Copy {} -->> {}'.format(file[0], newPhotoPath))
        shutil.copy(file[0], newPhotoPath)
        index += 1
