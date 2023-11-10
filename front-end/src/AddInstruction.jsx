import { useState, useRef } from "react";

const hostUrl = "http:\/\/localhost:3000/upload";

export const AddInstruction = () => {
    const filePicker = useRef(null);
    const [selectedFile, setSelectedFile] = useState(null);
    const [uploaded, setUploaded] = useState();

    const handleChange = (event) => {
        setSelectedFile(event.target.files)
    };


    const handleUpload = async () => {
        if (!selectedFile) {
            alert("Please select a file");
            return;
        }

        const title = document.getElementById("Title").value;
        const description = document.getElementById("Description").value;

        const formData = new FormData();

        for (let i = 0; ; i++) {
            if (selectedFile[i] === undefined) {
                break
            }
            formData.append('file', selectedFile[i]);
        }

        const res = await fetch(hostUrl, {
            method: 'POST',
            body: formData,
        });
        const data = await res.json();

        const payload = {
            action: "add_instruction",
            instruction: {
                title: title,
                description: description,
                path: data.filePath
            }
        }

        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const body = {
            method: 'POST',
            body: JSON.stringify(payload),
            headers: headers,
        }

        fetch("http:\/\/localhost:8080/handle", body)
            .then((response) => response.json())
            .then((data) => {
                if (data.error) {
                    console.log("error: ", data.message)
                } else {
                    console.log("Response from broker service: ", data.message);
                }
            })

        setUploaded(data);
    };

    const handlePick = () => {
        filePicker.current.click();
    }

    return (
        <>
            <div>
                <h1>Add Instruction</h1>
                <input
                    type="text"
                    placeholder="Type title"
                    id="Title"
                />
                <input
                    type="text"
                    placeholder="Type description"
                    id="Description"
                />
            </div>

            <button onClick={handlePick}>Pick files</button>
            <input
                className="hidden"
                type="file"
                ref = {filePicker}
                multiple
                onChange={handleChange}
                accept="image/*,.png,.jpg,.jpeg,.gif,.web,"
            />

            <button onClick={handleUpload}>Upload now</button>

            {selectedFile && (
                <h2>Instruction is ready to be add</h2>
            )}

            {uploaded && (
                <h2>"Success added"</h2>
            )}
        </>
    );
};