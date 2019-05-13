import * as mongoose from 'mongoose';

const Schema = mongoose.Schema;

export const FixtureSchema = new Schema({
    _id: {
        type: String,
        required: 'ID is required'
    },
    hTeam: {
        type: String,
        required: 'Home team is required'
    },
    aTeam: {
        type: String,
        required: 'Away team is required'
    },
    hGoals: {
        type: Number
    },
    aGoals: {
        type: Number
    },
    played: {
        type: Number            
    },
    dateTime: {
        type: Date,
        required: 'DateTime is required'
    },
    matchday: {
        type: Number,
        required: 'Matchday is required'
    }
});