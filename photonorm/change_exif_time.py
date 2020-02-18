#!/usr/bin/python
# -*- coding: UTF-8 -*-

import datetime
import sys
import pyexiv2


def ModifyImageFile(filename):
    image = pyexiv2.Image(filename)
    meta = image.read_exif()

    if 'Exif.Photo.DateTimeOriginal' in meta:
        date = meta['Exif.Photo.DateTimeOriginal']
        print (date)

        date_time_obj = datetime.datetime.strptime(date, '%Y:%m:%d %H:%M:%S')       
        date_time_obj = date_time_obj + datetime.timedelta(hours=3, minutes=0, seconds=0)
        print (date_time_obj)
        
        image.modify_exif({'Exif.Photo.DateTimeOriginal': date_time_obj.strftime('%Y:%m:%d %H:%M:%S')})
