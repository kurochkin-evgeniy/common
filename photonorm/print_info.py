#!/usr/bin/python
# -*- coding: UTF-8 -*-

import sys
import os
import datetime
import remove_dubles
import organize_folders

import pyexiv2


def IsImageFile(path):
    filename, ext = os.path.splitext(path)
    ext.upper()
    if ext == '.NEF' or ext == '.JPG' or ext == '.JPEG':
        return True
    return False


def PrintLine(label, value):
    message = u' {:<80}{}'.format(label, value)
    print(message)


def PrintExifTag(tag):
    value = ''
    label = tag.key

    try:
        if hasattr(tag, 'human_value'):
            value = tag.human_value
        elif hasattr(tag, 'raw_values'):
            if len(tag.raw_values) == 1:
                value = tag.raw_values[0]
            else:
                value = '!!!!!!!!!'
        else:
            if isinstance(tag.value, dict):
                value = tag.value['x-default'].encode('utf-8')
            elif isinstance(tag.value, list):
                value = tag.raw_value[0]
            else:
                value = tag.raw_value

        if hasattr(tag, 'label'):
            label = tag.label + '[' + tag.key + ']'

        value = value.decode('utf-8')
    except:
        value = "????"

    PrintLine(label, value)


def PrintExifAll(meta):

    print(' %-80s%s' % ("Comment", meta.comment))

    for key in meta.exif_keys:
        PrintExifTag(meta[key])

    print ("\n")
    for key in meta.iptc_keys:
        PrintExifTag(meta[key])

    print ("\n")
    for key in meta.xmp_keys:
        PrintExifTag(meta[key])


def PrintExifTagFromMeta(meta, tagName):
    try:
        PrintExifTag(meta[tagName])
    except:
        PrintLine(tagName, '\'\'')


def PrintExifMy(meta):

    # Genegic
    print(' %-80s%s' % ("Comment", meta.comment))
    PrintExifTagFromMeta(meta, 'Exif.Photo.DateTimeOriginal')
    PrintExifTagFromMeta(meta, 'Exif.Image.Model')

    # Caption
    PrintExifTagFromMeta(meta, 'Iptc.Application2.Caption')
    PrintExifTagFromMeta(meta, 'Xmp.dc.description')

    # Title
    PrintExifTagFromMeta(meta, 'Xmp.dc.title')
    PrintExifTagFromMeta(meta, 'Iptc.Application2.ObjectName')

    # Copyright
    PrintExifTagFromMeta(meta, 'Xmp.dc.rights')
    PrintExifTagFromMeta(meta, 'Iptc.Application2.Copyright')

    # CREATOR
    PrintExifTagFromMeta(meta, 'Xmp.dc.creator')
    PrintExifTagFromMeta(meta, 'Iptc.Application2.Byline')

    # HeadLine
    PrintExifTagFromMeta(meta, 'Iptc.Application2.Headline')
    PrintExifTagFromMeta(meta, 'Xmp.photoshop.Headline')

    PrintExifTagFromMeta(
        meta, 'Xmp.iptc.CreatorContactInfo/Iptc4xmpCore:CiAdrCity')
    PrintExifTagFromMeta(
        meta, 'Xmp.iptc.CreatorContactInfo/Iptc4xmpCore:CiAdrCtry')
    PrintExifTagFromMeta(
        meta, 'Xmp.iptc.CreatorContactInfo/Iptc4xmpCore:CiEmailWork')
    PrintExifTagFromMeta(
        meta, 'Xmp.iptc.CreatorContactInfo/Iptc4xmpCore:CiTelWork')
    PrintExifTagFromMeta(meta, 'Xmp.iptc.Location')

    # CountryName
    PrintExifTagFromMeta(meta, 'Xmp.photoshop.Country')
    PrintExifTagFromMeta(meta, 'Iptc.Application2.CountryName')

    # Credit
    PrintExifTagFromMeta(meta, 'Xmp.photoshop.Credit')
    PrintExifTagFromMeta(meta, 'Iptc.Application2.Credit')

    # Source
    PrintExifTagFromMeta(meta, 'Xmp.photoshop.Source')
    PrintExifTagFromMeta(meta, 'Iptc.Application2.Source')

    PrintExifTagFromMeta(meta, 'Xmp.xmpRights.UsageTerms')


def PrintExif(meta, all):
    if all:
        PrintExifAll(meta)
    else:
        PrintExifMy(meta)


def PrintExifForFile(path, all):
    print ("---------------------------------------------------")
    print (os.path.split(path)[1])
    meta = pyexiv2.ImageMetadata(path)
    meta.read()
    PrintExif(meta, all)


def PrintExifForDir(path, all):
    for root, dirs, files in os.walk(path):
        for _file in files:
            fullPath = os.path.join(root, _file)
            if IsImageFile(fullPath):
                PrintExifForFile(fullPath, all)


def PrintExifInfo(path, all):
    if os.path.isfile(path):
        PrintExifForFile(path, all)
    else:
        PrintExifForDir(path, all)
