#!/usr/bin/env python
# -*- coding: utf-8 -*-

from greet import greet as g


def test_greet():
    assert g() == "Hello, World!"
