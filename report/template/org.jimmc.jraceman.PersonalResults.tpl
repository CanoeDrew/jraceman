{{/*GT: {
  "display": "Personal Results",
  "description": "Personal results for each person collected on one page per person, sorted by team and last name",
  "where": [ "meet", "team", "person" ],
  "orderby": [
    {
      "name": "team",
      "display": "Team",
      "sql": "team.shortname"
    },
    {
      "name": "person",
      "display": "Person",
      "sql": "team.lastname, team.firstname"
    }
  ]
} */ -}}
{{ $RaceInfo := include "org.jimmc.jraceman.RaceInfo" -}}
{{ $comp := computed -}}
{{/* Use the selection criteria to get a list of person IDs.
      Use JOIN rather than LEFT JOIN on Entry so that
      we don't get anybody who has no entries. */ -}}
{{ $peopleIdRows := rows (printf `
    SELECT person.id as personId, Meet.id as meetId
    FROM person
      LEFT JOIN Team on person.teamId=Team.id
      JOIN Entry on person.id=Entry.personId
      LEFT JOIN Event on Entry.eventId=Event.id
      LEFT JOIN Meet on Event.meetId=Meet.id
    %s
    GROUP BY Person.id,Meet.id
    %s
` $comp.Where.WhereClause $comp.OrderBy.Clause) . -}}
{{ $sectionCountTable := `
	(SELECT Race.eventId AS eventId, Race.round AS round, count(*) AS sectionCount
	FROM Race JOIN Stage on Race.stageId=Stage.id
	GROUP BY Race.eventId,Race.round) as SectionsPerRound` -}}
<html>
<body>
{{ range $peopleIdRows -}}
{{ $personHeaderRow := row (printf `
    SELECT 
      Person.firstName as firstName, Person.lastName as lastName,
      Team.name as Team, Team.shortName as shortName,
      Meet.name as Meet
    FROM Person
      LEFT JOIN Team on Person.teamId=Team.id
      LEFT JOIN Entry on Person.id=Entry.personId
      LEFT JOIN Event on Entry.eventId=Event.id
      LEFT JOIN Meet on Event.meetId=Meet.id
    WHERE
      person.id = '%s' and meet.id = '%s'
    GROUP BY Person.id, Meet.id
    ` .personId .meetId) -}}
  <div class="header">
{{ include "org.jimmc.jraceman.datePrintedLine" }}
  </div>
  <center>
  <div class="titleArea">
  <h2>Results for {{$personHeaderRow.firstName}} {{$personHeaderRow.lastName}} of {{$personHeaderRow.Team}} ({{$personHeaderRow.shortName}})<br>{{$personHeaderRow.Meet}}</h2>
  </div>
  <table class="main" border=1>
    <tr class="rowHeader">
      <th>Event</th>
      <th>Race</th>
      <th>Scheduled Start</th>
      <th>Lane</th>
      <th>Result</th>
      <th>Place</th>
      <th>Score Place</th>
      <th>Score</th>
    </tr>
{{ $personResults := rows (printf `
    SELECT 
        (
	    CASE WHEN Event.scratched THEN '<strike>' ELSE '' END ||
	    '#' || Event.number || ' ' || Event.name ||
	    CASE WHEN Event.scratched THEN '</strike>' ELSE '' END
        ) as Event,
        (
	    CASE WHEN Race.scratched THEN '<strike>' ELSE '' END ||
	    '#' ||
            Race.number ||
	    ' ' ||Stage.name || ' ' ||
	    CASE WHEN SectionsPerRound.sectionCount>1 THEN 
	    	CASE WHEN Stage.isFinal THEN substring('ABCDEFGHIJK',Race.section,1) ELSE Race.section END ELSE '' END ||
	    CASE WHEN Race.scratched THEN '</strike>' ELSE '' END
        ) as Race,
        Race.scheduledStart as RaceTime,
	(
	    CASE WHEN Entry.scratched THEN '<strike>' ELSE '' END ||
	    Lane.lane ||
	    CASE WHEN Entry.scratched THEN '</strike>' ELSE '' END
	) as Lane,
        Lane.exceptionId,
        Exception.shortName as Exception,
        ROUND(Lane.result,3) as Result,
        Lane.place as Place,
	(
	    CASE WHEN Event.scratched OR Race.scratched OR Entry.scratched THEN '<strike>' ELSE '' END ||
	    Lane.scorePlace ||
	    CASE WHEN Event.scratched OR Race.scratched OR Entry.scratched THEN '</strike>' ELSE '' END
	) as ScorePlace,
	(
	    CASE WHEN Event.scratched OR Race.scratched OR Entry.scratched THEN '<strike>' ELSE '' END ||
	    Lane.score ||
	    CASE WHEN Event.scratched OR Race.scratched OR Entry.scratched THEN '</strike>' ELSE '' END
	) as Score
    FROM Person
      LEFT JOIN Team on Person.teamId=Team.id
      LEFT JOIN Entry on Person.id=Entry.personId
      LEFT JOIN Lane on Entry.id=Lane.entryId
      LEFT JOIN Race on Lane.raceId=Race.id
      LEFT JOIN Stage on Race.stageId=Stage.id
      LEFT JOIN Event on Entry.eventId=Event.id
      LEFT JOIN Exception on Lane.exceptionId=Exception.id
      LEFT JOIN Meet on Event.meetId=Meet.id
      LEFT JOIN %s
      	on (Race.round=SectionsPerRound.round and Race.eventId=SectionsPerRound.eventId)
    WHERE
      person.id = '%s' and meet.id = '%s'
      AND Lane.lane>=0 AND Race.scheduledStart is not null
    ORDER BY Race.scheduledStart`
    $sectionCountTable .personId .meetId) -}}
{{ range $personResults -}}
    <tr class="rowParity{rowParity}">
      <td>{{.Event}}</td>
      <td>{{.Race}}</td>
      <td>{{.RaceTime}}</td>
      <td>{{.Lane}}</td>
      <td>{{.Result}},{{.Exception}}</td>
      <td>{{.Place}}</td>
      <td>{{.ScorePlace}}</td>
      <td>{{.Score}}</td>
    </tr>
{{ end -}}
  </table>
{{ $personalTotal := row (printf `
    SELECT sum(Lane.score) as totalScore
    FROM Lane
      LEFT JOIN Entry on Lane.entryId=Entry.id
      LEFT JOIN Race on Lane.raceId=Race.id
      LEFT JOIN Person on Entry.personId=Person.id
      LEFT JOIN Event on Entry.eventId=Event.id
      LEFT JOIN Meet on Event.meetId=Meet.id
    WHERE
      person.id = '%s' and meet.id = '%s'
      AND NOT COALESCE(Event.scratched,false)
      AND NOT COALESCE(Race.scratched,false)
      AND NOT COALESCE(Entry.scratched,false)`
    .personId .meetId) -}}
        <div class="rowTotal">
	Total Points: {{$personalTotal.totalScore}}
        </div>
  <div class="footer">
{{ include "org.jimmc.jraceman.datePrintedLine" }}
  </div>
  </center>
{{ end -}}
</body>
</html>
