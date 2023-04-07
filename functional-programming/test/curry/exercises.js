import { filter, test } from 'ramda';
import QUnit from 'qunit';

const filterQs = filter(test(/q/ig));

QUnit.test("filter", assert => {
  assert.deepEqual(filterQs(['quick', 'camels', 'quarry', 'over', 'quails']), ['quick', 'quarry', 'quails']);
})
