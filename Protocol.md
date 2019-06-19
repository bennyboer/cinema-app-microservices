# Kommunikation der Services

## Grundlagen

Jeder Service, der mit anderen Services Kontakt aufnehmen will, erhält im Konstruktor ein `Dependencies` **struct**.
In diesem werden Funktionen (Supplier) übergeben, welche Instanzen von den konkreten Services zurückliefert.

### Beispiel

````go
// Dependencies the reservation service depends on.
type ReservationServiceDependencies struct {
	PresentationService func() presentation.PresentationService
	UserService         func() user.UserService
	CinemaService       func() cinema.CinemaService
}
````

So müssen die Services beim Start der Services noch nicht wissen, ob diese aktuell existieren, sondern erst wenn diese wirklich benötigt werden.
Es ist also egal, welcher Service zuerst gestartet wird.

### Proto

Die einzelnen Requests und Response Typen werden aus den anderen Services importiert und können für die Aufrufe verwendet werden.

````go
cinemaService := h.getCinemaService()

rsp, err := cinemaService.AreAvailable(context, &cinema.AvailableRequest{
    Id:    cinemaID,
    Seats: cinemaSeats,
})
````

## Auflistung der Abhängigkeiten

Im Folgenden werden die Abhängigkeiten der einzelnen Services untereinander aufgezeigt.

### Movie Service

- **Presentation Service**: Wenn ein Film gelöscht wird, wird der Presentation Service benachrichtigt und löscht alle damit verbundenen Präsentationen.

### User Service

- **Reservation Service**: Wenn ein Nutzer gelöscht wird, müssen auch alle mit ihm verbundenen Reservierungen gelöscht werden.

### Cinema Service

- **Presentation Service**: Wenn ein Kino gelöscht wird, wird der Presentation Service benachrichtigt. Alle mit dem Kino verbundenen Präsentationen müssen gelöscht werden.

### Presentation Service

- **Reservation Service**: Falls eine Präsentation gelöscht wird, wird der Reservation Service benachrichtigt, dass er die damit verbundenen Reservierungen kündigt.

### Reservation Service

- **Presentation Service**: Da eine Reservierung nur für eine Präsentation geht, muss über den Presentation Service die zugehörige Kino ID erfragt werden.
- **User Service**: Bevor eine Reservierung möglich ist, muss überprüft werden, ob die übergebene Benutzer ID einem existierenden Benutzer gehört.
- **Cinema Service**: Bei einer Reservierung für eine Präsentation in einem Kino, muss zuerst das dazugehörige Kino nach freien Plätzen befragt werden.
Außerdem werden bei erfolgreicher Reservierung die Sitze als belegt markiert.
