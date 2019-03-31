package report

import (
  "fmt"
  "testing"

  "github.com/google/go-cmp/cmp"
)

func TestClientVisibleReports(t *testing.T) {
  reports, err := ClientVisibleReportsOne("template")
  if err != nil {
    t.Fatalf("ClientVisibleReports error: %v", err)
  }
  expected := []*ReportFields{
    {
      Name: "org.jimmc.jraceman.Entries",
      Display: "Entries",
      Description: "Entries ordered as selected.",
      OrderBy: []FieldsOrderByItem{
        FieldsOrderByItem{Name: "team",        Display: "Team, Person, Event"},
        FieldsOrderByItem{Name: "person",      Display: "Person, Event"},
        FieldsOrderByItem{Name: "eventTeam",   Display: "Event, Team, Person"},
        FieldsOrderByItem{Name: "eventPerson", Display: "Event, Person"},
      },
    },
    {
      Name: "org.jimmc.jraceman.Lanes",
      Display: "Lanes",
      Description: "",
      OrderBy: []FieldsOrderByItem{},
    },
  }
  got, want := reports, expected
  if diff := cmp.Diff(want, got); diff != "" {
    t.Errorf("ClientVisibleReports() mismatch (-want +got):\n%s", diff)
  }
}

func TestReadTemplateAttrs(t *testing.T) {
  attrslist, err := ReadTemplateAttrs("template")
  if err != nil {
    t.Fatalf("ReadTemplateAttrs error: %v", err)
  }
  if got, want := len(attrslist), 2; got != want {
    t.Fatalf("Wrong number of files-with-attributes found, got %d, want %d", got, want)
  }
  fmt.Printf("attrslist: %+v\n", attrslist)
  if got, want := attrslist[0].Name, "org.jimmc.jraceman.Entries"; got != want {
    t.Errorf("Name of first report: got %s, want %s", got, want)
  }
}

func TestExtractUserOrderByMap(t *testing.T) {
  tests := []struct{
    name string
    input *ReportAttributes
    expect []FieldsOrderByItem
    expectError bool
  } {
    {
      name: "empty",
      input: &ReportAttributes{},
      expect: []FieldsOrderByItem{},
    },
    {
      name: "normal",
      input: &ReportAttributes{
        Name: "normal",
        OrderBy: []AttributesOrderByItem{
          {Name:"a", Display: "AA"},
          {Name:"b", Display: "BB"},
        },
      },
      expect: []FieldsOrderByItem{
        {Name: "a", Display: "AA"},
        {Name: "b", Display: "BB"},
      },
    },
  }
  for _, tc := range tests {
    t.Run(tc.name, func(t *testing.T) {
      got, err := extractUserOrderByList(tc.input)
      if tc.expectError {
        if err == nil {
          t.Fatalf("extractUserOrderByList: expected error but did not get one")
        }
      } else if err != nil {
        t.Fatalf("extractUserOrderByList: unexpected error: %v", err)
      } else {
        want := tc.expect
        if diff := cmp.Diff(want, got); diff != "" {
          t.Errorf("extractUserOrderByList mismatch (-want +got):\n%s", diff)
        }
      }
    })
  }
}
