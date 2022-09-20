{{/*GT: {
  "display": "Race Results",
  "description": "Results for a Race, sorted by finish order",
  "where": [ "event", "race" ],
  "orderby": [
    {
      "name": "event",
      "display": "Event, Race",
      "sql": "event.number, race.number"
    },
    {
      "name": "race",
      "display": "Race",
      "sql": "race.number"
    }
  ]
} */ -}}
{{ $RaceInfo := include "org.jimmc.jraceman.RaceInfo" -}}
{{ $comp := computed -}}
{{ $raceRows := rows (printf `
    SELECT %s
    FROM %s
    %s %s
` $RaceInfo.raceColumns $RaceInfo.raceTables $comp.Where.WhereClause $comp.OrderBy.Clause) . -}}
<html>
<body>
  <div class="header">
{{ include "org.jimmc.jraceman.datePrintedLine" }}
  </div>
  <center>
{{ range $raceRows -}}
  <div class="titleArea">
    <h3>{{.eventInfo}}, {{.raceNumberInfo}}</h3>
  </div>
  <table class="main" border=1>
    <tr class="rowHeader">
      <th width=60>Lane</th>
      <th width=300 alight=left>Person</th>
      <th width=100 alight=left>Team</th>
      <th width=80>Result</th>
      <th width=60>Place</th>
      <th width=60>Score Place</th>
    </tr>
{{ $laneRows := rows (printf `
    SELECT lane.lane as Lane,
        (person.firstname || ' ' || person.lastname) as Person,
        team.shortname as Team,
        ROUND(lane.result,3) as Result,
        lane.place as Place,
        lane.ScorePlace as ScorePlace
    FROM lane
    LEFT JOIN race on lane.raceid = race.id
    LEFT JOIN entry on lane.entryid = entry.id
    LEFT JOIN person on entry.personid = person.id
    LEFT JOIN team on person.teamid = team.id
    WHERE lane.raceid = '%s'
    ORDER BY Place` .raceid) -}}
{{ range $laneRows }}
    <tr class="rowParity{{ evenodd .rowindex "0" "1" }}">
      <td align=center>{{.Lane}}</td>
      <td>{{.Person}}</td>
      <td>{{.Team}}</td>
      <td align=right>{{.Result}}</td>
      <td align=center>{{.Place}}</td>
      <td align=center>{{.ScorePlace}}</td>
    </tr>
{{ end }}
  </table>
{{ end }}
  <div class="footer">
{{ include "org.jimmc.jraceman.datePrintedLine" }}
  </div>
</body>
</html>
