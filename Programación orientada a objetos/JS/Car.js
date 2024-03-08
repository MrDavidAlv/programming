function Car(licence, driver) {
    this.id;
    this.licence=licence;
    this.driver=driver;
    this.passanger;
}

Car.prototype.printDataCar = function(){
    console.log(this.driver);
    console.log(this.driver.name);
    console.log(this.driver.document);
}
