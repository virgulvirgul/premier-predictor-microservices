export class Forms {
    teams: Map<string, TeamForm>
}

export class TeamForm {
    forms: TeamMatchResult[]
}

export class TeamMatchResult {
    result: Result;
    score: string;
    opponent: string;
    location: Location;
}

export enum Result {
    WIN = 'win',
    DRAW = 'draw',
    LOSS = 'loss',
}

export enum Location {
    HOME = 'home',
    AWAY = 'away',
}