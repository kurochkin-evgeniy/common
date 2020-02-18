#!/usr/bin/python
# -*- coding: UTF-8 -*-

import sys
import os
import hashlib


def GetAllFilesWithSizes(path):
    resultFiles = []
    for root, dirs, files in os.walk(path):
        for _file in files:
            fullPath = os.path.join(root, _file)
            resultFiles.append((fullPath, os.path.getsize(fullPath)))
    return resultFiles


def Md5ForFile(fname):
    hash_md5 = hashlib.md5()
    with open(fname, "rb") as f:
        for chunk in iter(lambda: f.read(4096), b""):
            hash_md5.update(chunk)
    return hash_md5.hexdigest()


def FindDublicatesByMd5(paths):
    if len(paths) <= 1:
        return []

    dublicates = []
    uniqHashes = set()
    for file in paths:
        hash = Md5ForFile(file[0])
        if hash in uniqHashes:
            dublicates.append(file[0])
        else:
            uniqHashes.add(hash)

    return dublicates


def FindDublicatedFiles(path):
    files = GetAllFilesWithSizes(path)
    files.sort(key=lambda tup: tup[1])

    if len(files) == 0:
        return

    startIndex = 0
    currentSize = files[0][1]
    dublicates = []

    for index in range(0, len(files)):
        if currentSize != files[index][1]:
            d = FindDublicatesByMd5(files[startIndex: index])
            if d:
                dublicates.extend(d)
            startIndex = index
            currentSize = files[index][1]

    d = FindDublicatesByMd5(files[startIndex:])
    if d:
         dublicates.extend(d)
    return dublicates


def DeleteFiles(paths):
    for file in paths:
        os.remove(file)


def RemoveDublicatedFiles(path, emulate):
    dublicates = FindDublicatedFiles(path)
    if not emulate:
        DeleteFiles(dublicates)
    return dublicates
