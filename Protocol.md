# Kommunikation der Services

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