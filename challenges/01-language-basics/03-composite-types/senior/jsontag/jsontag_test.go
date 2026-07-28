package jsontag

import "testing"

func TestMarshal(t *testing.T) {
	got, err := Marshal(User{FirstName: "Ada", LastName: "Lovelace"})
	if err != nil {
		t.Fatal(err)
	}
	want := `{"first_name":"Ada","last_name":"Lovelace"}`
	if got != want {
		t.Errorf("Marshal=%s; want %s", got, want)
	}
}
