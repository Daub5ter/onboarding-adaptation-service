import {useState} from "react";

export const GetAllInstructions = () => {
    const [showed, setShowed] = useState();

    const handleShow = async () => {
        const payload = {
            action: "get_all_instructions",
        }

        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const response = await fetch("http:\/\/localhost:8080/handle", {
            method: 'POST',
            body: JSON.stringify(payload),
            headers: headers,
        });

        const data = await response.json();

        setShowed(data);
    }


    return (
        <>
            <button onClick={handleShow}>show all instruction</button>

            {showed && (
                <div>
                    <h2>test</h2>
                    <h2>{showed.data[3].title}</h2>
                    <h2>{showed.data[3].description}</h2>
                </div>

            )}
        </>
    )

}