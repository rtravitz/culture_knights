
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE table meetings
(
  id Serial,
  bookId INT,
  jesus TEXT,
  jesusChooserId INT,
  jesusExplanation TEXT,
  bookChooserId INT,
  goldenExplanation TEXT,
  goldenRecipientId INT,
  leadExplanation TEXT,
  leadRecipientId INT,
  discussionDate TIMESTAMP
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE meetings;

