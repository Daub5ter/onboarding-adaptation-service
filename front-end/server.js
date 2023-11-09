const cors = require('cors');
const express = require('express');
const fileUpload = require('express-fileupload');

const port = 3000;

const app = express();

app.use(fileUpload({
    createParentPath: true,
}))

app.use(cors())

app.post('/upload', (req, res) => {


    if (!req.files) {
        console.log("no file uploaded")
        return res.status(400).json({ msg: 'No file uploaded' });
    }

    const file = req.files.file;

    if (!file) {
        console.log("Incorrect input name")
        return res.json({ error: 'Incorrect input name' });
    }

    console.log("Creating mv")

    file.mv(`${__dirname}/images/${file.name}`, err => {
        if (err) {
            console.error(err);
            return res.status(500).send(err);
        }
        console.log('file was uploaded');

        res.json({
            fileName: file.name,
            filePath: `/images/${file.name}`,
        });
    });
});

app.listen(port, () => console.log(`Server Started in port ${port}`));