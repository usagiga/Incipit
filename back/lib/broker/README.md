# Broker

The broker who establishes connection to DB.
This library is compatible with `github.com/jinzhu/gorm`.

## Overview

In Docker or the other virtual environment,
running Go app / initializing DB spends bit time.
Although, unfortunatelly,
`gorm.Open()` try to connect only once.
So, you have to wait for launching DB to connect correctly.

Broker waits for it.
