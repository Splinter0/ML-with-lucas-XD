import numpy, operator

##classifying romantic and action movies giving values
##as coordinates in a 2D space
def createDataSet():
    group = numpy.array([[1,2],[1,4],[4,1],[4,2]])
    labels = ["Action", "Action", "Romantic", "Romantic"]
    return group, labels

class KNN(object):
    def __init__(self, k=3):
        self.k = k

    def getData(self, data, labels):
        self.data = data
        self.labels = labels

    def classify(self, x):
        sortedDistance = self.euclideanDistance(x)
        majority = self.majorityClass(sortedDistance)
        return majority[0][0]

    def euclideanDistance(self, x):
        size = self.data.shape[0]
        # numpy.tile(A, reps): construct an array by repeating A
        # the number given by reps
        # do the differece of all points from x
        diffMat = numpy.tile(x, (size, 1)) - self.data
        # distance is equal to the square root of the sum of the squares of the coordinates
        d = (diffMat ** 2).sum(axis=1) ** 0.5
        # sort indices and return them
        return d.argsort()

    def majorityClass(self, sortedDistance):
        countLabels = {}
        # iterate for how many times k
        for i in range(self.k):
            label = self.labels[sortedDistance[i]]
            countLabels[label] = countLabels.get(label, 0) +1
        # count how many different labes in the sortedDistance in the first k indices
        # sort the count of the labels in descending order
        return sorted(countLabels.iteritems(), operator.itemgetter(1), reverse=True)


group, labels = createDataSet()
a = KNN()
a.getData(group, labels)
result = a.classify([2,3])
print(result)
