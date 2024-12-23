import ollama from 'ollama'
import fs from 'fs'

import {simpleParser } from 'mailparser'

import { HTMLToJSON } from 'html-to-json-parser'; 

const filename = process.argv[2];
console.log('received command line args', process.argv);


const file = fs.readFileSync(filename, 'utf-8')

const parsed = await simpleParser(file)

const data = parsed.text

const toJSON = await HTMLToJSON(data, true)
console.log("JSON", toJSON)
console.log('=================================================================')

const response = await ollama.chat({
  model: 'llama3.2',
  messages: [
      { role: 'user', content: data + ' \n above is a email, can you tell me if this email is a reservation email for a flight, hotel or a resturant?' }

  ],
})
console.log(response.message.content)

