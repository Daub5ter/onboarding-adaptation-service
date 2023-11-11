import { Onboarding } from './Onboarding'
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
import {Adaptation} from "./Adaptation";

function App() {
    return (
        <div className="App">
            <Onboarding />
            <br/>
            <Adaptation />
        </div>
    );
}

export default App