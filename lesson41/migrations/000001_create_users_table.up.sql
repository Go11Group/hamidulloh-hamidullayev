create type passing as enum ('pass', 'fail');
create type appilation as enum ('impossible');

create table exam_result
(
    id       uuid primary key,
    name     varchar not null,
    is_pass  passing    not null,
    score    int8 default 59 check (score < 100 and score >= 0),
    feedback text default 'yiqildingiz, yaxshilab o''qing, balki dasturlash sizga emas',
    shikoyat appilation not null
);
