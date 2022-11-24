const fs = require('fs');

let tmp_file = "/tmp/temporary_file"; // Sensitive
fs.readFile(tmp_file, 'utf8', function (err, data) {
  // ...
});
