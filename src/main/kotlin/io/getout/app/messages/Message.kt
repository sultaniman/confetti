package io.getout.app.messages

import org.hibernate.annotations.CreationTimestamp
import java.time.Instant
import java.util.*
import javax.persistence.*

@Entity
@Table(name = "messages")
data class Message(
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(name = "id", insertable = false, updatable = false, nullable = false)
    val ID: UUID,

    @Column(name = "body")
    var Body: String = "",

    @CreationTimestamp
    @Column(name = "created_at")
    val CreatedAt: Instant? = null,

    @CreationTimestamp
    @Column(name = "updated_at")
    val UpdatedAt: Instant? = null,
)
