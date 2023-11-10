import { useState, useRef } from "react";

const hostUrl = "http:\/\/localhost:3000/upload";
let val = 1;

export const UploadFile = () => {
    const filePicker = useRef(null);
    const [selectedFile, setSelectedFile] = useState(null);
    const [uploaded, setUploaded] = useState();

    const handleChange = (event) => {
        val = document.getElementById('countImages').value;
        setSelectedFile(event.target.files)
    };


    const handleUpload = async () => {
        if (!selectedFile) {
            alert("Please select a file");
            return;
        }

        const formData = new FormData();

        console.log(selectedFile)

        for (let i = 0; i < val; i++) {
            formData.append('file', selectedFile[i]);
        }

        //console.log(formData);

        const res = await fetch(hostUrl, {
            method: 'POST',
            body: formData,
        });
        const data = await res.json();

        setUploaded(data);
    };

    const handlePick = () => {
        filePicker.current.click();
    }

    return (
        <>
            <label>Count of images</label>
            <input
                id="countImages"
                type="text"
            />

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
                <ul>
                    <li>Name: {selectedFile.name}</li>
                    <li>Type: {selectedFile.type}</li>
                    <li>Size: {selectedFile.size}</li>
                </ul>
            )}

            {uploaded && (
                <div>
                    <h2>{uploaded.filePath}</h2>
                    <img alt='' src={uploaded.filePath} width="200"/>
                </div>
            )}
        </>
    );
};