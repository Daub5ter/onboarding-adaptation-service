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
        </div>
    );
}

export default App