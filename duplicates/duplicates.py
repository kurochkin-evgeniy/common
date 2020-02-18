#!/usr/bin/python
# -*- coding: UTF-8 -*-

import sys
import os
import fnmatch
import hashlib
import time
from concurrent.futures import ThreadPoolExecutor
import fastio

class DupSearchParams:
    path = ''
    filter = ''

class DupSearchResult:
    totalSize = 0
    totalCount = 0
    duplicates = None

def GetAllFilesWithSizes(param):
    resultFiles = []
    
    for root, _, files in fastio.walk(param.path):
        for file in files:
            if fnmatch.fnmatch(file, param.filter):
                fullPath = os.path.join(root, file)
                stat = os.stat(fullPath)
                resultFiles.append((fullPath, stat.st_size, stat.st_ctime))

    print("os.walk complete!")  
    return resultFiles


def Md5ForFile(fname):
    hash_md5 = hashlib.md5()
    with open(fname, "rb") as f:
        for chunk in iter(lambda: f.read(1024*1024), b""):
            hash_md5.update(chunk)
    return hash_md5.hexdigest()


def FindDuplicatesByMd5InEqualSizeFiles(paths):
    if len(paths) <= 1:
        return []

    duplicates = []
    uniqHashes = {}
    for file in paths:
        hash = Md5ForFile(file[0])
        if hash in uniqHashes:
            duplicates.append((file[0], file[1], hash, uniqHashes[hash]))
        else:
            uniqHashes.update({hash: file[0]})

    duplicates.sort(key=lambda tup: tup[3])
    return duplicates


def FindDuplicatesByMd5InEqualSizeFilesTask(paths, result):
    d = FindDuplicatesByMd5InEqualSizeFiles(paths)
    if d:
        result.extend(d)  

def FindDuplicatedFiles(param):
    files = GetAllFilesWithSizes(param)
    
    files.sort(key=lambda tup: (tup[1], tup[2]))

    if len(files) == 0:
        return

    startIndex = 0
    currentSize = files[0][1]
    duplicates = []

    with ThreadPoolExecutor(max_workers=20) as executor:
        for index in range(0, len(files)):
            if currentSize != files[index][1]:
                executor.submit(FindDuplicatesByMd5InEqualSizeFilesTask, files[startIndex: index], duplicates)
                startIndex = index
                currentSize = files[index][1]

        executor.submit(FindDuplicatesByMd5InEqualSizeFilesTask, files[startIndex:], duplicates)


    result = DupSearchResult()
    result.totalCount = len(files) 
    result.totalSize = sum([f[1] for f in files])
    result.duplicates =  duplicates

    return result


def DuplicatedFiles(param):
    result = FindDuplicatedFiles(param)
    totalDubSize = sum([d[1] for d in result.duplicates])

    print("Full number of files = ", result.totalCount)
    print("Total size = ", result.totalSize)   
    print("Number of duplicates = " , len(result.duplicates))
    print("Total duplicates size = ", totalDubSize)
    print("Percent = ", totalDubSize/result.totalSize * 100)
    
    f = open("result.txt","w", encoding="utf-8")
    header = ''
    for file in result.duplicates: 
        if header != file[3]:
            header = file[3]
            f.write("\n%s (size = %d, hash = %s)\n" % (header, file[1], file[2]))  
        f.write("        %s\n" %  file[0])    

    f.close() 



if __name__ == '__main__':
    start_time = time.time()

    param = DupSearchParams()
    param.path   = sys.argv[1]
    param.filter = '*.pdf'

    DuplicatedFiles(param)
    print("--- %s seconds ---" % (time.time() - start_time))

