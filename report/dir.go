package report

import (
  "fmt"
  "log"

  "github.com/jimmc/gtrepgen/gen"
)

// ReportAttributes contains the attributes loaded from a report template.
type ReportAttributes struct {
  Name string
  Display string
  Description string
  Where []string
  OrderBy []AttributesOrderByItem
}

// AttributesOrderByItem contains the OrderBy info loaded from a report template.
type AttributesOrderByItem struct {
  Name string
  Display string
  Sql string
}

// ReportFields contains the information about a report template that is given to the user.
type ReportFields struct {
  Name string
  Display string
  Description string
  OrderBy []FieldsOrderByItem
}

// FieldsOrderByItem is the information about OrderBy that is given to the user.
type FieldsOrderByItem struct {
  Name string
  Display string
}

/* ClientVisibleReports returns the list of reports and their attributes
 * that should be visible to a client.
 * Once we have user ids and authorizations, this function should accept
 * a user id and return only data that should be visible to that user.
 */
func ClientVisibleReports(reportRoots []string) ([]*ReportFields, error) {
  allFields := make([]*ReportFields, 0)
  for _, root := range reportRoots {
    fields, err := ClientVisibleReportsOne(root)
    if err != nil {
      return nil, err
    }
    allFields = append(allFields, fields...)
  }
  return allFields, nil
}

/* ClientVisibleReportsOne returns the list of reports and their user-visible attributes
 * from one root directory.
 */
func ClientVisibleReportsOne(templateDir string) ([]*ReportFields, error) {
  attrs, err := ReadTemplateAttrs(templateDir)
  if err != nil {
    return nil, err
  }
  reports := make([]*ReportFields, 0)
  for _, tplAttrs := range attrs {
    userOrderByList, err := extractUserOrderByList(tplAttrs)
    if err != nil {
      // If we get an error, don't add this report to the list, but still show other reports.
      log.Printf("Error decoding orderby in template %q: %v", tplAttrs.Name, err)
      continue
    }
    if tplAttrs.Display == "" {
      // Don't include templates with no Display attribute.
      continue
    }
    report := &ReportFields{
      Name: tplAttrs.Name,
      Display: tplAttrs.Display,
      Description: tplAttrs.Description,
      OrderBy: userOrderByList,
    }
    reports = append(reports, report)
  }
  return reports, nil
}

/* ReadTemplateAttrs loads the attributes from all the template files in
 * the given directory.
 */
func ReadTemplateAttrs(templateDir string) ([]*ReportAttributes, error) {
  newDestPointer := func() interface{} {
    return &ReportAttributes{}
  }
  fileAttrs, err := gen.ReadDirFilesAttributesAs(templateDir, newDestPointer)
  if fileAttrs == nil {
    return nil, err
  }
  reportAttrs := []*ReportAttributes{}
  for _, fattrs := range fileAttrs {
    if fattrs.Err != nil {
      return nil, fmt.Errorf("for template %q received error %v", fattrs.Name, fattrs.Err)
    }
    attrs, ok := fattrs.Attributes.(*ReportAttributes)
    if !ok {
      return nil, fmt.Errorf("invalid data type for template %q", fattrs.Name)
    }
    attrs.Name = fattrs.Name
    reportAttrs = append(reportAttrs, attrs)
  }
  return reportAttrs, err
}

// extractUserOrderByList looks at the orderby attribute in the given template attributes
// and extacts from that the user-visible fields.
func extractUserOrderByList(tplAttrs *ReportAttributes) ([]FieldsOrderByItem, error) {
    r := []FieldsOrderByItem{}
    for _, v := range tplAttrs.OrderBy {
      r = append(r, FieldsOrderByItem{
        Name: v.Name,
        Display: v.Display,
      })
    }
    return r, nil
}
