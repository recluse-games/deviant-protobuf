// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: Entity.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Deviant {

  /// <summary>Holder for reflection information generated from Entity.proto</summary>
  public static partial class EntityReflection {

    #region Descriptor
    /// <summary>File descriptor for Entity.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static EntityReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CgxFbnRpdHkucHJvdG8SB0RldmlhbnQaD0FsaWdubWVudC5wcm90bxoQQXR0",
            "YWNobWVudC5wcm90bxoNQ2xhc3Nlcy5wcm90bxoQQ29uZGl0aW9ucy5wcm90",
            "bxoKSGFuZC5wcm90bxoKRGVjay5wcm90bxoNRGlzY2FyZC5wcm90bxoWRW50",
            "aXR5U3RhdGVOYW1lcy5wcm90bxoZRW50aXR5Um90YXRpb25OYW1lcy5wcm90",
            "byLPAwoGRW50aXR5EgoKAmlkGAEgASgJEgoKAmhwGAIgASgFEgoKAmFwGAMg",
            "ASgFEigKBXN0YXRlGAQgASgOMhkuRGV2aWFudC5FbnRpdHlTdGF0ZU5hbWVz",
            "EiUKCWFsaWdubWVudBgFIAEoDjISLkRldmlhbnQuQWxpZ25tZW50Eh8KBWNs",
            "YXNzGAYgASgOMhAuRGV2aWFudC5DbGFzc2VzEicKCmNvbmRpdGlvbnMYByAD",
            "KA4yEy5EZXZpYW50LkNvbmRpdGlvbnMSKAoLYXR0YWNobWVudHMYCCADKA4y",
            "Ey5EZXZpYW50LkF0dGFjaG1lbnQSGwoEaGFuZBgJIAEoCzINLkRldmlhbnQu",
            "SGFuZBIbCgRkZWNrGAogASgLMg0uRGV2aWFudC5EZWNrEiEKB2Rpc2NhcmQY",
            "CyABKAsyEC5EZXZpYW50LkRpc2NhcmQSEgoKaW5pdGlhdGl2ZRgMIAEoBRIP",
            "Cgdvd25lcklkGA0gASgJEg0KBW1heEhwGA4gASgFEg0KBW1heEFwGA8gASgF",
            "EgwKBG5hbWUYECABKAkSLgoIcm90YXRpb24YESABKA4yHC5EZXZpYW50LkVu",
            "dGl0eVJvdGF0aW9uTmFtZXNCC1oJLjtkZXZpYW50YgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Deviant.AlignmentReflection.Descriptor, global::Deviant.AttachmentReflection.Descriptor, global::Deviant.ClassesReflection.Descriptor, global::Deviant.ConditionsReflection.Descriptor, global::Deviant.HandReflection.Descriptor, global::Deviant.DeckReflection.Descriptor, global::Deviant.DiscardReflection.Descriptor, global::Deviant.EntityStateNamesReflection.Descriptor, global::Deviant.EntityRotationNamesReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Deviant.Entity), global::Deviant.Entity.Parser, new[]{ "Id", "Hp", "Ap", "State", "Alignment", "Class", "Conditions", "Attachments", "Hand", "Deck", "Discard", "Initiative", "OwnerId", "MaxHp", "MaxAp", "Name", "Rotation" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class Entity : pb::IMessage<Entity> {
    private static readonly pb::MessageParser<Entity> _parser = new pb::MessageParser<Entity>(() => new Entity());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<Entity> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Deviant.EntityReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Entity() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Entity(Entity other) : this() {
      id_ = other.id_;
      hp_ = other.hp_;
      ap_ = other.ap_;
      state_ = other.state_;
      alignment_ = other.alignment_;
      class_ = other.class_;
      conditions_ = other.conditions_.Clone();
      attachments_ = other.attachments_.Clone();
      hand_ = other.hand_ != null ? other.hand_.Clone() : null;
      deck_ = other.deck_ != null ? other.deck_.Clone() : null;
      discard_ = other.discard_ != null ? other.discard_.Clone() : null;
      initiative_ = other.initiative_;
      ownerId_ = other.ownerId_;
      maxHp_ = other.maxHp_;
      maxAp_ = other.maxAp_;
      name_ = other.name_;
      rotation_ = other.rotation_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Entity Clone() {
      return new Entity(this);
    }

    /// <summary>Field number for the "id" field.</summary>
    public const int IdFieldNumber = 1;
    private string id_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Id {
      get { return id_; }
      set {
        id_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "hp" field.</summary>
    public const int HpFieldNumber = 2;
    private int hp_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int Hp {
      get { return hp_; }
      set {
        hp_ = value;
      }
    }

    /// <summary>Field number for the "ap" field.</summary>
    public const int ApFieldNumber = 3;
    private int ap_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int Ap {
      get { return ap_; }
      set {
        ap_ = value;
      }
    }

    /// <summary>Field number for the "state" field.</summary>
    public const int StateFieldNumber = 4;
    private global::Deviant.EntityStateNames state_ = global::Deviant.EntityStateNames.Idle;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Deviant.EntityStateNames State {
      get { return state_; }
      set {
        state_ = value;
      }
    }

    /// <summary>Field number for the "alignment" field.</summary>
    public const int AlignmentFieldNumber = 5;
    private global::Deviant.Alignment alignment_ = global::Deviant.Alignment.Friendly;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Deviant.Alignment Alignment {
      get { return alignment_; }
      set {
        alignment_ = value;
      }
    }

    /// <summary>Field number for the "class" field.</summary>
    public const int ClassFieldNumber = 6;
    private global::Deviant.Classes class_ = global::Deviant.Classes.Warrior;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Deviant.Classes Class {
      get { return class_; }
      set {
        class_ = value;
      }
    }

    /// <summary>Field number for the "conditions" field.</summary>
    public const int ConditionsFieldNumber = 7;
    private static readonly pb::FieldCodec<global::Deviant.Conditions> _repeated_conditions_codec
        = pb::FieldCodec.ForEnum(58, x => (int) x, x => (global::Deviant.Conditions) x);
    private readonly pbc::RepeatedField<global::Deviant.Conditions> conditions_ = new pbc::RepeatedField<global::Deviant.Conditions>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::Deviant.Conditions> Conditions {
      get { return conditions_; }
    }

    /// <summary>Field number for the "attachments" field.</summary>
    public const int AttachmentsFieldNumber = 8;
    private static readonly pb::FieldCodec<global::Deviant.Attachment> _repeated_attachments_codec
        = pb::FieldCodec.ForEnum(66, x => (int) x, x => (global::Deviant.Attachment) x);
    private readonly pbc::RepeatedField<global::Deviant.Attachment> attachments_ = new pbc::RepeatedField<global::Deviant.Attachment>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<global::Deviant.Attachment> Attachments {
      get { return attachments_; }
    }

    /// <summary>Field number for the "hand" field.</summary>
    public const int HandFieldNumber = 9;
    private global::Deviant.Hand hand_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Deviant.Hand Hand {
      get { return hand_; }
      set {
        hand_ = value;
      }
    }

    /// <summary>Field number for the "deck" field.</summary>
    public const int DeckFieldNumber = 10;
    private global::Deviant.Deck deck_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Deviant.Deck Deck {
      get { return deck_; }
      set {
        deck_ = value;
      }
    }

    /// <summary>Field number for the "discard" field.</summary>
    public const int DiscardFieldNumber = 11;
    private global::Deviant.Discard discard_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Deviant.Discard Discard {
      get { return discard_; }
      set {
        discard_ = value;
      }
    }

    /// <summary>Field number for the "initiative" field.</summary>
    public const int InitiativeFieldNumber = 12;
    private int initiative_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int Initiative {
      get { return initiative_; }
      set {
        initiative_ = value;
      }
    }

    /// <summary>Field number for the "ownerId" field.</summary>
    public const int OwnerIdFieldNumber = 13;
    private string ownerId_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string OwnerId {
      get { return ownerId_; }
      set {
        ownerId_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "maxHp" field.</summary>
    public const int MaxHpFieldNumber = 14;
    private int maxHp_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int MaxHp {
      get { return maxHp_; }
      set {
        maxHp_ = value;
      }
    }

    /// <summary>Field number for the "maxAp" field.</summary>
    public const int MaxApFieldNumber = 15;
    private int maxAp_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int MaxAp {
      get { return maxAp_; }
      set {
        maxAp_ = value;
      }
    }

    /// <summary>Field number for the "name" field.</summary>
    public const int NameFieldNumber = 16;
    private string name_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Name {
      get { return name_; }
      set {
        name_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "rotation" field.</summary>
    public const int RotationFieldNumber = 17;
    private global::Deviant.EntityRotationNames rotation_ = global::Deviant.EntityRotationNames.North;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Deviant.EntityRotationNames Rotation {
      get { return rotation_; }
      set {
        rotation_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as Entity);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(Entity other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Id != other.Id) return false;
      if (Hp != other.Hp) return false;
      if (Ap != other.Ap) return false;
      if (State != other.State) return false;
      if (Alignment != other.Alignment) return false;
      if (Class != other.Class) return false;
      if(!conditions_.Equals(other.conditions_)) return false;
      if(!attachments_.Equals(other.attachments_)) return false;
      if (!object.Equals(Hand, other.Hand)) return false;
      if (!object.Equals(Deck, other.Deck)) return false;
      if (!object.Equals(Discard, other.Discard)) return false;
      if (Initiative != other.Initiative) return false;
      if (OwnerId != other.OwnerId) return false;
      if (MaxHp != other.MaxHp) return false;
      if (MaxAp != other.MaxAp) return false;
      if (Name != other.Name) return false;
      if (Rotation != other.Rotation) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Id.Length != 0) hash ^= Id.GetHashCode();
      if (Hp != 0) hash ^= Hp.GetHashCode();
      if (Ap != 0) hash ^= Ap.GetHashCode();
      if (State != global::Deviant.EntityStateNames.Idle) hash ^= State.GetHashCode();
      if (Alignment != global::Deviant.Alignment.Friendly) hash ^= Alignment.GetHashCode();
      if (Class != global::Deviant.Classes.Warrior) hash ^= Class.GetHashCode();
      hash ^= conditions_.GetHashCode();
      hash ^= attachments_.GetHashCode();
      if (hand_ != null) hash ^= Hand.GetHashCode();
      if (deck_ != null) hash ^= Deck.GetHashCode();
      if (discard_ != null) hash ^= Discard.GetHashCode();
      if (Initiative != 0) hash ^= Initiative.GetHashCode();
      if (OwnerId.Length != 0) hash ^= OwnerId.GetHashCode();
      if (MaxHp != 0) hash ^= MaxHp.GetHashCode();
      if (MaxAp != 0) hash ^= MaxAp.GetHashCode();
      if (Name.Length != 0) hash ^= Name.GetHashCode();
      if (Rotation != global::Deviant.EntityRotationNames.North) hash ^= Rotation.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Id.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(Id);
      }
      if (Hp != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(Hp);
      }
      if (Ap != 0) {
        output.WriteRawTag(24);
        output.WriteInt32(Ap);
      }
      if (State != global::Deviant.EntityStateNames.Idle) {
        output.WriteRawTag(32);
        output.WriteEnum((int) State);
      }
      if (Alignment != global::Deviant.Alignment.Friendly) {
        output.WriteRawTag(40);
        output.WriteEnum((int) Alignment);
      }
      if (Class != global::Deviant.Classes.Warrior) {
        output.WriteRawTag(48);
        output.WriteEnum((int) Class);
      }
      conditions_.WriteTo(output, _repeated_conditions_codec);
      attachments_.WriteTo(output, _repeated_attachments_codec);
      if (hand_ != null) {
        output.WriteRawTag(74);
        output.WriteMessage(Hand);
      }
      if (deck_ != null) {
        output.WriteRawTag(82);
        output.WriteMessage(Deck);
      }
      if (discard_ != null) {
        output.WriteRawTag(90);
        output.WriteMessage(Discard);
      }
      if (Initiative != 0) {
        output.WriteRawTag(96);
        output.WriteInt32(Initiative);
      }
      if (OwnerId.Length != 0) {
        output.WriteRawTag(106);
        output.WriteString(OwnerId);
      }
      if (MaxHp != 0) {
        output.WriteRawTag(112);
        output.WriteInt32(MaxHp);
      }
      if (MaxAp != 0) {
        output.WriteRawTag(120);
        output.WriteInt32(MaxAp);
      }
      if (Name.Length != 0) {
        output.WriteRawTag(130, 1);
        output.WriteString(Name);
      }
      if (Rotation != global::Deviant.EntityRotationNames.North) {
        output.WriteRawTag(136, 1);
        output.WriteEnum((int) Rotation);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Id.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Id);
      }
      if (Hp != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Hp);
      }
      if (Ap != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Ap);
      }
      if (State != global::Deviant.EntityStateNames.Idle) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) State);
      }
      if (Alignment != global::Deviant.Alignment.Friendly) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) Alignment);
      }
      if (Class != global::Deviant.Classes.Warrior) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) Class);
      }
      size += conditions_.CalculateSize(_repeated_conditions_codec);
      size += attachments_.CalculateSize(_repeated_attachments_codec);
      if (hand_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Hand);
      }
      if (deck_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Deck);
      }
      if (discard_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Discard);
      }
      if (Initiative != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Initiative);
      }
      if (OwnerId.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(OwnerId);
      }
      if (MaxHp != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(MaxHp);
      }
      if (MaxAp != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(MaxAp);
      }
      if (Name.Length != 0) {
        size += 2 + pb::CodedOutputStream.ComputeStringSize(Name);
      }
      if (Rotation != global::Deviant.EntityRotationNames.North) {
        size += 2 + pb::CodedOutputStream.ComputeEnumSize((int) Rotation);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(Entity other) {
      if (other == null) {
        return;
      }
      if (other.Id.Length != 0) {
        Id = other.Id;
      }
      if (other.Hp != 0) {
        Hp = other.Hp;
      }
      if (other.Ap != 0) {
        Ap = other.Ap;
      }
      if (other.State != global::Deviant.EntityStateNames.Idle) {
        State = other.State;
      }
      if (other.Alignment != global::Deviant.Alignment.Friendly) {
        Alignment = other.Alignment;
      }
      if (other.Class != global::Deviant.Classes.Warrior) {
        Class = other.Class;
      }
      conditions_.Add(other.conditions_);
      attachments_.Add(other.attachments_);
      if (other.hand_ != null) {
        if (hand_ == null) {
          Hand = new global::Deviant.Hand();
        }
        Hand.MergeFrom(other.Hand);
      }
      if (other.deck_ != null) {
        if (deck_ == null) {
          Deck = new global::Deviant.Deck();
        }
        Deck.MergeFrom(other.Deck);
      }
      if (other.discard_ != null) {
        if (discard_ == null) {
          Discard = new global::Deviant.Discard();
        }
        Discard.MergeFrom(other.Discard);
      }
      if (other.Initiative != 0) {
        Initiative = other.Initiative;
      }
      if (other.OwnerId.Length != 0) {
        OwnerId = other.OwnerId;
      }
      if (other.MaxHp != 0) {
        MaxHp = other.MaxHp;
      }
      if (other.MaxAp != 0) {
        MaxAp = other.MaxAp;
      }
      if (other.Name.Length != 0) {
        Name = other.Name;
      }
      if (other.Rotation != global::Deviant.EntityRotationNames.North) {
        Rotation = other.Rotation;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            Id = input.ReadString();
            break;
          }
          case 16: {
            Hp = input.ReadInt32();
            break;
          }
          case 24: {
            Ap = input.ReadInt32();
            break;
          }
          case 32: {
            State = (global::Deviant.EntityStateNames) input.ReadEnum();
            break;
          }
          case 40: {
            Alignment = (global::Deviant.Alignment) input.ReadEnum();
            break;
          }
          case 48: {
            Class = (global::Deviant.Classes) input.ReadEnum();
            break;
          }
          case 58:
          case 56: {
            conditions_.AddEntriesFrom(input, _repeated_conditions_codec);
            break;
          }
          case 66:
          case 64: {
            attachments_.AddEntriesFrom(input, _repeated_attachments_codec);
            break;
          }
          case 74: {
            if (hand_ == null) {
              Hand = new global::Deviant.Hand();
            }
            input.ReadMessage(Hand);
            break;
          }
          case 82: {
            if (deck_ == null) {
              Deck = new global::Deviant.Deck();
            }
            input.ReadMessage(Deck);
            break;
          }
          case 90: {
            if (discard_ == null) {
              Discard = new global::Deviant.Discard();
            }
            input.ReadMessage(Discard);
            break;
          }
          case 96: {
            Initiative = input.ReadInt32();
            break;
          }
          case 106: {
            OwnerId = input.ReadString();
            break;
          }
          case 112: {
            MaxHp = input.ReadInt32();
            break;
          }
          case 120: {
            MaxAp = input.ReadInt32();
            break;
          }
          case 130: {
            Name = input.ReadString();
            break;
          }
          case 136: {
            Rotation = (global::Deviant.EntityRotationNames) input.ReadEnum();
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code