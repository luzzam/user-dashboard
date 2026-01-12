const fs = require('fs');
const path = require('path');

class Parser {
  constructor(filePath) {
    this.filePath = filePath;
    this.data = [];
  }

  readData() {
    try {
      const rawData = fs.readFileSync(this.filePath, 'utf8');
      const lines = rawData.split('\n');
      lines.forEach((line) => {
        const parsedLine = this.parseLine(line);
        if (parsedLine) {
          this.data.push(parsedLine);
        }
      });
    } catch (error) {
      throw new Error(`Error reading file: ${error.message}`);
    }
  }

  parseLine(line) {
    const regex = /^(\d{4}-\d{2}-\d{2})\s+(\d{2}:\d{2}:\d{2})\s+(.*)$/;
    const match = line.match(regex);
    if (match) {
      return {
        date: match[1],
        time: match[2],
        message: match[3].trim()
      };
    } else {
      return null;
    }
  }

  getData() {
    return this.data;
  }
}

module.exports = Parser;