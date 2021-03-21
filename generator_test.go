package generator

import (
	"io/ioutil"
	"testing"
)

func TestNew(t *testing.T) {
	doc, _ := New(Invoice, &Options{
		TextTypeInvoice: "INVOICE",
		AutoPrint:       true,
	})

	doc.Options.CurrencySymbol = "$"
	doc.Options.TextItemsUnitCostTitle = "Production Fee"
	doc.Options.TextItemsQuantityTitle = "Shipping"
	doc.Options.TextItemsNameTitle = "Order Number"

	doc.SetHeader(&HeaderFooter{
		Text:       "<center>Corvallis3D LLC Invoice for Craftcloud by All3DP.</center>",
		Pagination: true,
	})

	doc.SetFooter(&HeaderFooter{
		Text:       "<center>Corvallis3D LLC Invoice for Craftcloud by All3DP.</center>",
		Pagination: true,
	})

	doc.SetRef("1")
	doc.SetVersion("1.0")

	doc.SetDescription("A description")
	doc.SetNotes("I love croissant cotton candy. Carrot cake sweet I love sweet roll cake powder! I love croissant cotton candy. Carrot cake sweet I love sweet roll cake powder! I love croissant cotton candy. Carrot cake sweet I love sweet roll cake powder! I love croissant cotton candy. Carrot cake sweet I love sweet roll cake powder! ")

	doc.SetDate("02/03/2021")
	doc.SetPaymentTerm("02/04/2021")

	logoBytes, _ := ioutil.ReadFile("invoice_logo.png")

	doc.SetCompany(&Contact{
		Name: "Corvallis3D LLC",
		Logo: &logoBytes,
		Address: &Address{
			Address:    "520 NW Oak Ave.",
			Address2:   "Ste B",
			PostalCode: "97330",
			City:       "Corvallis, OR",
			Country:    "United States of America",
		},
	})

	doc.SetCustomer(&Contact{
		Name: "All3DP GmbH",
		Address: &Address{
			Address:    "Ridlerstr. 31A",
			PostalCode: "80339",
			City:       "Munich",
			Country:    "Germany",
		},
	})

	for i := 0; i < 2; i++ {
		doc.AppendItem(&Item{
			Name:         "Cupcake ipsum dolor sit amet bonbon, coucou bonbon lala jojo, mama titi toto",
			Description:  "Cupcake ipsum dolor sit amet bonbon, Cupcake ipsum dolor sit amet bonbon, Cupcake ipsum dolor sit amet bonbon",
			UnitCost:     "99876.89",
			ShippingCost: "9.99",
			Quantity:     "1",
			Discount: &Discount{
				Percent: "12.5",
			},
		})
	}

	doc.AppendItem(&Item{
		Name:         "Test",
		UnitCost:     "99876.89",
		ShippingCost: "9.99",
		Quantity:     "1",
		Discount: &Discount{
			Percent: "12.5",
		},
	})

	doc.AppendItem(&Item{
		Name:         "Test",
		UnitCost:     "3576.89",
		ShippingCost: "9.99",
		Quantity:     "1",
		Discount: &Discount{
			Percent: "12.5",
		},
	})

	doc.AppendItem(&Item{
		Name:         "Test",
		UnitCost:     "889.89",
		ShippingCost: "9.99",
		Quantity:     "1",
		Discount: &Discount{
			Percent: "12.5",
		},
	})

	pdf, err := doc.Build()
	if err != nil {
		t.Errorf(err.Error())
	}

	err = pdf.OutputFileAndClose("out.pdf")

	if err != nil {
		t.Errorf(err.Error())
	}
}
