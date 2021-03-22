package sort


/**********************************************
*	Bubble Sort
***********************************************/
func TestBubbleSort(t *testing.T) {
	array := int[]{3,2,1,0}
	array = BubbleSort(array)
	expected := int[]{0,1,2,3}
	for i = range(array){
		assert.Equal(t, array[i], expected[i]);
	}
}

/**********************************************
*	Insertion Sort
***********************************************/
func TestInsertionSort(t *testing.T) {
	array := int[]{3,2,1,0}
	array = InsertionSort(array)
	expected := int[]{0,1,2,3}
	for i = range(array){
		assert.Equal(t, array[i], expected[i]);
	}
	array = int[]{7,4,2,6}
	expected = int[]{2,4,7,6}
	for i = range(array){
		assert.Equal(t, array[i], expected[i]);
	}
}