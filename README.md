# Multiplex Kino Applikation mit Microservice Architektur

## Installationshinweise

> Stellen Sie sicher, dass Sie [Go](https://golang.org/) (> **v1.12**) installiert haben

Führen Sie im Wurzelverzeichnis des Repositories `./build.sh` aus. 
Sie erhalten in jedem Unterverzeichnis der Services (`cinema`, `user`, `movie`, `reservation`, `presentation`) ein Binary zum Ausführen.

## Integrationstests

Im Unterverzeichnis `integrationtests` befinden sich Integrationstests für die Microservices.
Dort wird unter anderem getestet...

- ...ob beim Löschen eines Kinos auch die zugehörigen Vorstellungen und Reservierungen gelöscht werden.
- ...was bei einer gleichzeitigen Reservierung von überschneidenden Sitzplätzen von zwei unterschiedlichen Nutzern auf die selbe Vorstellung passiert (Nur ein Nutzer sollte die Reservierung durchführen können).

#### Ausführen

Die Integrationstests sind normale Go "Unit Tests" im Unterverzeichnis `integrationtests`.
Zum Ausführen befolgen Sie folgende Schritte:

1. Starten Sie **alle Services**
1. Navigieren Sie in den Ordner `integrationtests`
1. Führen Sie Folgendes aus:
   ```bash
   go test github.com/ob-vss-ss19/blatt-4-sudo_blatt4/integrationtests
   ```
1. Falls die Tests funktioniert haben sollte Folgendes zu sehen sein:
   ```bash
    ok      github.com/ob-vss-ss19/blatt-4-sudo_blatt4/integrationtests     1.447s
   ```