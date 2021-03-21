package sort


/**********************************************
*	BubbleSort
***********************************************/
func TestBubbleSort(t *testing.T) {
	array := int[]{3,2,1,0}
	array = BubbleSort(array)
	expected := int[]{0,1,2,3}
	for i = range(array){
		assert.Equal(t, array[i], expected[i]);
	}
}