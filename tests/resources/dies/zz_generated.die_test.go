//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by diegen. DO NOT EDIT.

package dies

import (
	testing "dies.dev/testing"
	testingx "testing"
)

func TestClusterRunTemplateDie_MissingMethods(t *testingx.T) {
	die := ClusterRunTemplateBlank
	ignore := []string{"TypeMeta", "ObjectMeta"}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ClusterRunTemplateDie: %s", diff.List())
	}
}

func TestRunTemplateSpecDie_MissingMethods(t *testingx.T) {
	die := RunTemplateSpecBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for RunTemplateSpecDie: %s", diff.List())
	}
}

func TestRunnableDie_MissingMethods(t *testingx.T) {
	die := RunnableBlank
	ignore := []string{"TypeMeta", "ObjectMeta"}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for RunnableDie: %s", diff.List())
	}
}

func TestRunnableSpecDie_MissingMethods(t *testingx.T) {
	die := RunnableSpecBlank
	ignore := []string{"Inputs"}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for RunnableSpecDie: %s", diff.List())
	}
}

func TestRunnableStatusDie_MissingMethods(t *testingx.T) {
	die := RunnableStatusBlank
	ignore := []string{"Outputs"}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for RunnableStatusDie: %s", diff.List())
	}
}
