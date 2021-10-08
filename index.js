const { Socket } = require('dgram');
const express = require('express');
const path = require('path')

require('dotenv').config();


const app = express()
const http = require('http').createServer(app)

app.use(express.static(path.join(__dirname,'public')))

app.get('/forms', (req, res) => {
    res.sendFile(__dirname + '/form.html');
  });

const io = require('socket.io')(http)
io.on('connection',Socket=>{
    console.log('connection ready')
})

const PORT = process.env.PORT || 8080 
http.listen(PORT,()=>{
    console.log('server is runing on', PORT);
})