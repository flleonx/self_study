class Boat {
  constructor (hasMotor, motorCount, motorBrand, motorModel, 
               hasSails, sailsCount, sailsMaterial, sailsColor,
               hullColor, hasCabin) {
    // ...
  }
}

/*
  Invoking such a constructor would create some hard to read code,
  which is easily prone to errors (which argument is what?).
*/

const myBoat = new Boat(true, 2, 'Best Motor Co.', 'OM123', true, 1,
                        'fabric', 'white', 'blue', false);
