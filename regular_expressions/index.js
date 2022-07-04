const str = "goooooeeelll";
const regex = /go*/;

const result = str.match(regex);

const username = "AY231";
const userCheck = /^[A-Za-z]{2,}\d*$/;
const checked = userCheck.test(username);

console.log('CHECKED', checked);
