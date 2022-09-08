package dict

import "testing"

func TestDelete(t *testing.T) {
    
    d := Dictionary{"test":"test_val"}
    //err := d.Update("test", "updated_val")

    d.Delete("test")

    _, err := d.Search("test")

    if err != ErrNotFound {
        t.Errorf("expected %q to be deleted", "test")
    }
}

func TestUpdate(t *testing.T) {

    t.Run("existing word", func(t *testing.T) {

        d := Dictionary{"test":"test_val"}
        err := d.Update("test", "updated_val")

        assertError(t, err, nil)
        assertDefinition(t, d, "test", "updated_val")
    })

    t.Run("new  word", func(t *testing.T) {

        d := Dictionary{}
        err := d.Update("test", "updated_val")

        assertError(t, err, ErrKeyDoesNotExist)
        //assertDefinition(t, d, "test", "updated_val")
    })
}

func TestAdd(t *testing.T) {
    t.Run("new word", func(t *testing.T) {
        d := Dictionary{}
        err := d.Add("test", "test_val")
        
        assertError(t, err, nil)
        assertDefinition(t, d, "test", "test_val")
    })

    t.Run("same word", func(t *testing.T) {
        d := Dictionary{ "test": "test_val" }
        err := d.Add("test", "new_val")
        
        assertError(t, err, ErrExists)
        assertDefinition(t, d, "test", "test_val")
    })
}

func assertDefinition(t testing.TB, d Dictionary, word, val string) {
    t.Helper()

    got, err := d.Search(word)

    //want := "test_val"

    if err != nil {
        t.Fatal("got err while adding:", err)
    }

    if got != val {
        t.Errorf("got %q want %q", got, val)
    }
} 

func TestSearch(t *testing.T) {
    
    //dictionary := map[string]string{ "test": "this is a test" }

    dictionary := Dictionary{ "test": "this is a test" }

    t.Run("Know key", func(t *testing.T){

        got, _ := dictionary.Search("test")

        want := "this is a test"

        assertStrings(t, got, want)
    })

    t.Run("unknown key", func(t *testing.T) {
        
        _, err := dictionary.Search("unknown")

        assertError(t, err, ErrNotFound)
    })
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q, test %q", got, want, "test")
    }
}

func assertError(t testing.TB, got, want error) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
