-- name: GetSportsmenBySportID :many
SELECT sm.*
FROM sportsman_sports ss
JOIN sportsmen sm ON sm.id = ss.sportsman_id
WHERE
	ss.sport_id = @sport_id
	AND (@rank::SMALLINT IS NULL OR ss.rank = @rank);
