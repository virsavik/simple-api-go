// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testInvoiceProducts(t *testing.T) {
	t.Parallel()

	query := InvoiceProducts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testInvoiceProductsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceProduct{}
	if err = randomize.Struct(seed, o, invoiceProductDBTypes, true, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := InvoiceProducts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInvoiceProductsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceProduct{}
	if err = randomize.Struct(seed, o, invoiceProductDBTypes, true, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := InvoiceProducts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := InvoiceProducts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInvoiceProductsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceProduct{}
	if err = randomize.Struct(seed, o, invoiceProductDBTypes, true, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := InvoiceProductSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := InvoiceProducts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInvoiceProductsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceProduct{}
	if err = randomize.Struct(seed, o, invoiceProductDBTypes, true, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := InvoiceProductExists(ctx, tx, o.InvoiceID, o.ProductID)
	if err != nil {
		t.Errorf("Unable to check if InvoiceProduct exists: %s", err)
	}
	if !e {
		t.Errorf("Expected InvoiceProductExists to return true, but got false.")
	}
}

func testInvoiceProductsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceProduct{}
	if err = randomize.Struct(seed, o, invoiceProductDBTypes, true, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	invoiceProductFound, err := FindInvoiceProduct(ctx, tx, o.InvoiceID, o.ProductID)
	if err != nil {
		t.Error(err)
	}

	if invoiceProductFound == nil {
		t.Error("want a record, got nil")
	}
}

func testInvoiceProductsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceProduct{}
	if err = randomize.Struct(seed, o, invoiceProductDBTypes, true, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = InvoiceProducts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testInvoiceProductsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceProduct{}
	if err = randomize.Struct(seed, o, invoiceProductDBTypes, true, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := InvoiceProducts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testInvoiceProductsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	invoiceProductOne := &InvoiceProduct{}
	invoiceProductTwo := &InvoiceProduct{}
	if err = randomize.Struct(seed, invoiceProductOne, invoiceProductDBTypes, false, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}
	if err = randomize.Struct(seed, invoiceProductTwo, invoiceProductDBTypes, false, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = invoiceProductOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = invoiceProductTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := InvoiceProducts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testInvoiceProductsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	invoiceProductOne := &InvoiceProduct{}
	invoiceProductTwo := &InvoiceProduct{}
	if err = randomize.Struct(seed, invoiceProductOne, invoiceProductDBTypes, false, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}
	if err = randomize.Struct(seed, invoiceProductTwo, invoiceProductDBTypes, false, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = invoiceProductOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = invoiceProductTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := InvoiceProducts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func invoiceProductBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceProduct) error {
	*o = InvoiceProduct{}
	return nil
}

func invoiceProductAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceProduct) error {
	*o = InvoiceProduct{}
	return nil
}

func invoiceProductAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceProduct) error {
	*o = InvoiceProduct{}
	return nil
}

func invoiceProductBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceProduct) error {
	*o = InvoiceProduct{}
	return nil
}

func invoiceProductAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceProduct) error {
	*o = InvoiceProduct{}
	return nil
}

func invoiceProductBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceProduct) error {
	*o = InvoiceProduct{}
	return nil
}

func invoiceProductAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceProduct) error {
	*o = InvoiceProduct{}
	return nil
}

func invoiceProductBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceProduct) error {
	*o = InvoiceProduct{}
	return nil
}

func invoiceProductAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceProduct) error {
	*o = InvoiceProduct{}
	return nil
}

func testInvoiceProductsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &InvoiceProduct{}
	o := &InvoiceProduct{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, invoiceProductDBTypes, false); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct object: %s", err)
	}

	AddInvoiceProductHook(boil.BeforeInsertHook, invoiceProductBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	invoiceProductBeforeInsertHooks = []InvoiceProductHook{}

	AddInvoiceProductHook(boil.AfterInsertHook, invoiceProductAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	invoiceProductAfterInsertHooks = []InvoiceProductHook{}

	AddInvoiceProductHook(boil.AfterSelectHook, invoiceProductAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	invoiceProductAfterSelectHooks = []InvoiceProductHook{}

	AddInvoiceProductHook(boil.BeforeUpdateHook, invoiceProductBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	invoiceProductBeforeUpdateHooks = []InvoiceProductHook{}

	AddInvoiceProductHook(boil.AfterUpdateHook, invoiceProductAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	invoiceProductAfterUpdateHooks = []InvoiceProductHook{}

	AddInvoiceProductHook(boil.BeforeDeleteHook, invoiceProductBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	invoiceProductBeforeDeleteHooks = []InvoiceProductHook{}

	AddInvoiceProductHook(boil.AfterDeleteHook, invoiceProductAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	invoiceProductAfterDeleteHooks = []InvoiceProductHook{}

	AddInvoiceProductHook(boil.BeforeUpsertHook, invoiceProductBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	invoiceProductBeforeUpsertHooks = []InvoiceProductHook{}

	AddInvoiceProductHook(boil.AfterUpsertHook, invoiceProductAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	invoiceProductAfterUpsertHooks = []InvoiceProductHook{}
}

func testInvoiceProductsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceProduct{}
	if err = randomize.Struct(seed, o, invoiceProductDBTypes, true, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := InvoiceProducts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testInvoiceProductsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceProduct{}
	if err = randomize.Struct(seed, o, invoiceProductDBTypes, true); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(invoiceProductColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := InvoiceProducts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testInvoiceProductToOneInvoiceUsingInvoice(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local InvoiceProduct
	var foreign Invoice

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, invoiceProductDBTypes, false, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, invoiceDBTypes, false, invoiceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Invoice struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.InvoiceID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Invoice().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := InvoiceProductSlice{&local}
	if err = local.L.LoadInvoice(ctx, tx, false, (*[]*InvoiceProduct)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Invoice == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Invoice = nil
	if err = local.L.LoadInvoice(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Invoice == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testInvoiceProductToOneProductUsingProduct(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local InvoiceProduct
	var foreign Product

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, invoiceProductDBTypes, false, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, productDBTypes, false, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ProductID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Product().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := InvoiceProductSlice{&local}
	if err = local.L.LoadProduct(ctx, tx, false, (*[]*InvoiceProduct)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Product == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Product = nil
	if err = local.L.LoadProduct(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Product == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testInvoiceProductToOneSetOpInvoiceUsingInvoice(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a InvoiceProduct
	var b, c Invoice

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, invoiceProductDBTypes, false, strmangle.SetComplement(invoiceProductPrimaryKeyColumns, invoiceProductColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, invoiceDBTypes, false, strmangle.SetComplement(invoicePrimaryKeyColumns, invoiceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, invoiceDBTypes, false, strmangle.SetComplement(invoicePrimaryKeyColumns, invoiceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Invoice{&b, &c} {
		err = a.SetInvoice(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Invoice != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.InvoiceProducts[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.InvoiceID != x.ID {
			t.Error("foreign key was wrong value", a.InvoiceID)
		}

		if exists, err := InvoiceProductExists(ctx, tx, a.InvoiceID, a.ProductID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testInvoiceProductToOneSetOpProductUsingProduct(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a InvoiceProduct
	var b, c Product

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, invoiceProductDBTypes, false, strmangle.SetComplement(invoiceProductPrimaryKeyColumns, invoiceProductColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, productDBTypes, false, strmangle.SetComplement(productPrimaryKeyColumns, productColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, productDBTypes, false, strmangle.SetComplement(productPrimaryKeyColumns, productColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Product{&b, &c} {
		err = a.SetProduct(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Product != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.InvoiceProducts[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ProductID != x.ID {
			t.Error("foreign key was wrong value", a.ProductID)
		}

		if exists, err := InvoiceProductExists(ctx, tx, a.InvoiceID, a.ProductID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testInvoiceProductsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceProduct{}
	if err = randomize.Struct(seed, o, invoiceProductDBTypes, true, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testInvoiceProductsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceProduct{}
	if err = randomize.Struct(seed, o, invoiceProductDBTypes, true, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := InvoiceProductSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testInvoiceProductsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceProduct{}
	if err = randomize.Struct(seed, o, invoiceProductDBTypes, true, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := InvoiceProducts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	invoiceProductDBTypes = map[string]string{`InvoiceID`: `uuid`, `ProductID`: `uuid`, `Amount`: `integer`, `UnitPrice`: `numeric`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                     = bytes.MinRead
)

func testInvoiceProductsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(invoiceProductPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(invoiceProductAllColumns) == len(invoiceProductPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceProduct{}
	if err = randomize.Struct(seed, o, invoiceProductDBTypes, true, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := InvoiceProducts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, invoiceProductDBTypes, true, invoiceProductPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testInvoiceProductsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(invoiceProductAllColumns) == len(invoiceProductPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceProduct{}
	if err = randomize.Struct(seed, o, invoiceProductDBTypes, true, invoiceProductColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := InvoiceProducts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, invoiceProductDBTypes, true, invoiceProductPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(invoiceProductAllColumns, invoiceProductPrimaryKeyColumns) {
		fields = invoiceProductAllColumns
	} else {
		fields = strmangle.SetComplement(
			invoiceProductAllColumns,
			invoiceProductPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := InvoiceProductSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testInvoiceProductsUpsert(t *testing.T) {
	t.Parallel()

	if len(invoiceProductAllColumns) == len(invoiceProductPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := InvoiceProduct{}
	if err = randomize.Struct(seed, &o, invoiceProductDBTypes, true); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert InvoiceProduct: %s", err)
	}

	count, err := InvoiceProducts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, invoiceProductDBTypes, false, invoiceProductPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize InvoiceProduct struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert InvoiceProduct: %s", err)
	}

	count, err = InvoiceProducts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
