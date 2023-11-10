import { AddInstruction } from './AddInstruction'
import { AddUsersInstruction} from "./AddUsersInstruction";
import { GetAllInstructions } from './GetAllInstructions'
import { GetOneInstruction } from "./GetOneInstruction";

function App() {
    return (
        <div className="App">
            <GetAllInstructions />
            <GetOneInstruction />
            <AddUsersInstruction />
            <AddInstruction />
        </div>
    );
}

export default App