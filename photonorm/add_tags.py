#!/usr/bin/python
# -*- coding: UTF-8 -*-

import datetime
import os
import sys
import json
import pyexiv2

def IsImageFile(path):
    _, ext = os.path.splitext(path)
    ext.upper()
    if ext == '.NEF' or ext == '.JPG' or ext == '.JPEG':
        return True
    return False


def GetUserMetaData():
    with open('meta.json') as f:
      return json.load(f)


def AddExif(path):
    meta = pyexiv2.ImageMetadata(path)
    meta.read()

    userData = GetUserMetaData()

    Description = userData['event']['description']
    City = userData['event']['city']
    Country = userData['event']['country']
    Location = userData['event']['location']

    Author = userData['author']['author']
    Phone = userData['author']['phone']
    Email = userData['author']['email']
    AuthorCity = userData['author']['authorCity']
    AuthorCountry = userData['author']['authorCountry']

    # destination
    meta['Iptc.Application2.Caption'] = [Description]
    meta['Xmp.dc.description'] = Description
    meta['Iptc.Application2.Headline'] = [Description]
    meta['Xmp.photoshop.Headline'] = Description
    meta['Xmp.dc.title'] = Description
    meta['Iptc.Application2.ObjectName'] = [Description]

    # Author
    meta['Xmp.dc.rights'] = Author
    meta['Iptc.Application2.Copyright'] = [Author]
    meta['Xmp.dc.creator'] = [Author]
    meta['Iptc.Application2.Byline'] = [Author]

    meta['Xmp.iptc.CreatorContactInfo/Iptc4xmpCore:CiTelWork'] = Phone
    meta['Xmp.iptc.CreatorContactInfo/Iptc4xmpCore:CiEmailWork'] = Email
    meta['Xmp.iptc.CreatorContactInfo/Iptc4xmpCore:CiAdrCity'] = AuthorCity
    meta['Xmp.iptc.CreatorContactInfo/Iptc4xmpCore:CiAdrCtry'] = AuthorCountry

    meta['Xmp.iptc.Location'] = Location

    meta['Xmp.photoshop.Country'] = Country
    meta['Iptc.Application2.CountryName'] = [Country]

    meta['Xmp.photoshop.City'] = City
    meta['Iptc.Application2.City'] = [City]

    meta.write()


def AddTags(path):
    for root, _, files in os.walk(path):
        for _file in files:
            fullPath = os.path.join(root, _file)
            if IsImageFile(fullPath):
                AddExif(fullPath)
