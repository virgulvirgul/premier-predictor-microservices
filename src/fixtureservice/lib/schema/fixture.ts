import * as mongoose from 'mongoose';

const Schema = mongoose.Schema;

export const FixtureSchema = new Schema({
    _id: {
        type: String,
        required: 'Enter ID'
    },
    hTeam: {
        type: String,
        required: 'Enter a home team'
    },
    aTeam: {
        type: String,
        required: 'Enter a away team'
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
    },
    matchday: {
        type: Number
    }
});