
class UberPool extends Car {
	String brand;
	String model;
	
	public UberPool(String licence, Account driver, String brand, String model){
		super(licence, driver);
		this.brand = brand;
		this.model = model;
		
	}
}
