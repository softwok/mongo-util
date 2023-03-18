package builder

import (
	f "github.com/softwok/mongo-util/field"
	o "github.com/softwok/mongo-util/operator"
	"go.mongodb.org/mongo-driver/bson"
)

// Bucket function returns a mongo $bucket operator used in aggregations.
func Bucket(groupBy, boundaries, def, output interface{}) Operator {
	m := bson.M{}

	appendNotNull(m, f.GroupBy, groupBy)
	appendNotNull(m, f.Boundaries, boundaries)
	appendNotNull(m, f.Default, def)
	appendNotNull(m, f.Output, output)

	return New(o.Bucket, m)
}

// BucketAuto function returns a mongo $bucketAuto operator used in aggregations.
func BucketAuto(groupBy, buckets, output, granularity interface{}) Operator {
	m := bson.M{}

	appendNotNull(m, f.GroupBy, groupBy)
	appendNotNull(m, f.Buckets, buckets)
	appendNotNull(m, f.Output, output)
	appendNotNull(m, f.Granularity, granularity)

	return New(o.BucketAuto, m)
}

// CollStats function returns a mongo $collStats operator used in aggregations.
func CollStats(latencyStats, storageStats, count interface{}) Operator {
	m := bson.M{}

	appendNotNull(m, f.LatencyStats, latencyStats)
	appendNotNull(m, f.StorageStats, storageStats)
	appendNotNull(m, f.Count, count)

	return New(o.CollStats, m)
}

// CurrentOp function returns a mongo $currentOp operator used in aggregations.
func CurrentOp(allUsers, idleConnections, idleCursors, idleSessions, localOps interface{}) Operator {
	m := bson.M{}

	appendNotNull(m, f.AllUsers, allUsers)
	appendNotNull(m, f.IdleConnections, idleConnections)
	appendNotNull(m, f.IdleCursors, idleCursors)
	appendNotNull(m, f.IdleSessions, idleSessions)
	appendNotNull(m, f.LocalOps, localOps)

	return New(o.CurrentOp, m)
}

// $geoNear,$graphLookup has many params, those functions
// will have too many params and do not make readable code.

// Group function returns a mongo $group operator used in aggregations.
func Group(ID interface{}, params bson.M) Operator {
	m := bson.M{}

	appendNotNull(m, f.ID, ID)

	for key, val := range params {
		appendNotNull(m, key, val)
	}

	return New(o.Group, m)
}

// Lookup function returns a mongo $lookup operator used in aggregations.
func Lookup(from, localField, foreignField, as interface{}) Operator {
	m := bson.M{}

	appendNotNull(m, f.From, from)
	appendNotNull(m, f.LocalField, localField)
	appendNotNull(m, f.ForeignField, foreignField)
	appendNotNull(m, f.As, as)

	return New(o.Lookup, m)
}

// UncorrelatedLookup function returns a mongo $lookup operator used in aggregations.
func UncorrelatedLookup(from, let, pipeline, as interface{}) Operator {
	m := bson.M{}

	appendNotNull(m, f.From, from)
	appendNotNull(m, f.Let, let)
	appendNotNull(m, f.Pipeline, pipeline)
	appendNotNull(m, f.As, as)

	return New(o.Lookup, m)
}

// Merge function returns a mongo $merge operator used in aggregations.
func Merge(into, on, let, whenMatched, whenNotMatched interface{}) Operator {
	m := bson.M{}

	appendNotNull(m, f.Into, into)
	appendNotNull(m, f.On, on)
	appendNotNull(m, f.Let, let)
	appendNotNull(m, f.WhenMatched, whenMatched)
	appendNotNull(m, f.WhenNotMatched, whenNotMatched)

	return New(o.Merge, m)
}

// ReplaceRoot function returns a mongo $replaceRoot operator used in aggregations.
func ReplaceRoot(newRoot interface{}) Operator {
	m := bson.M{}

	appendNotNull(m, f.NewRoot, newRoot)

	return New(o.ReplaceRoot, m)
}

// Sample function returns a mongo sample operator used in aggregations.
func Sample(size interface{}) Operator {
	m := bson.M{}

	appendNotNull(m, f.Size, size)

	return New(o.Sample, m)
}

// Unwind function returns a mongo $unwind operator used in aggregations.
func Unwind(path, includeArrayIndex, preserveNullAndEmptyArrays interface{}) Operator {
	m := bson.M{}

	appendNotNull(m, f.Path, path)
	appendNotNull(m, f.IncludeArrayIndex, includeArrayIndex)
	appendNotNull(m, f.PreserveNullAndEmptyArrays, preserveNullAndEmptyArrays)

	return New(o.Unwind, m)
}
