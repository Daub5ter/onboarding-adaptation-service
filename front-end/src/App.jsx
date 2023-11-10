import { AddInstruction } from './AddInstruction'
import { AddUsersInstruction} from "./AddUsersInstruction";
import { SolveInstruction } from "./SolveInstruction";
import { GetAllInstructions } from './GetAllInstructions'
import { GetOneInstruction } from "./GetOneInstruction";
import { GetUsersInstructions } from "./GetUsersInstructions";
import { GetPercentInstructions } from "./GetPercentInstructions";
import {GetAllUsers} from "./GetAllUsers";
import {GetUserByID} from "./GetUserByID";
import {AuthUser} from "./AuthUser";
import {GetPercentKnowledge} from "./GetPercentKnowledge";
import {AddKnowledge} from "./AddKnowledge";
import {GetAllKnowledge} from "./GetAllKnowledge";
import {AddUsersKnowledge} from "./AddUsersKnowledge";

function App() {
    return (
        <div className="App">
            <GetAllInstructions />
            <GetOneInstruction />
            <AddUsersInstruction />
            <SolveInstruction />
            <GetUsersInstructions />
            <GetPercentInstructions />
            <AddInstruction />

            <GetAllUsers />
            <GetUserByID />
            <AuthUser />

            <br></br>

            <AddKnowledge />
            <GetAllKnowledge />
            <AddUsersKnowledge />
            <GetPercentKnowledge />
        </div>
    );
}

export default App