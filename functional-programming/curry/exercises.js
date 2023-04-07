import { filter, test } from 'ramda';
import QUnit from 'qunit';

const filterQs = filter(test(/q/ig));

const filteredWords = filterQs(['quick', 'camels', 'quarry', 'over', 'quails']);

console.log(filteredWords);
